package usecase

import (
	"fmt"
	"time"

	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/helper"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
)

func (u OrderUsecase) OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface) {
	order, terr = u.orderRepo.OrderList(ctx)

	return
}

func (u OrderUsecase) OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface) {
	order, terr = u.orderRepo.OrderGetByID(ctx, id)

	return
}

func (u OrderUsecase) OrderItem(ctx *gin.Context, req []models.Item) (paymentOrders models.PaymentWithOrder, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.orderRepo.GetDB)
	defer func() {
		if terr != nil {
			helper.TxRollBack(ctx)
		} else {
			helper.TxCommit(ctx)
		}
	}()

	var (
		products   = make(map[int64]models.Product)
		totalPrice int64
	)

	// totaling price
	reqMap := make(map[int64]int64)
	for _, v := range req {
		var pd models.Product
		pd, terr = u.orderRepo.ProductGetByID(ctx, v.ProductID)
		if terr != nil {
			return
		}

		products[int64(pd.ID)] = pd

		totalPrice += (v.Amount * pd.Price)

		if rm, ok := reqMap[v.ProductID]; !ok {
			reqMap[v.ProductID] = v.Amount
		} else {
			reqMap[v.ProductID] = rm + v.Amount
		}
	}

	// create Payments
	code := utstring.RandomString(20)
	for {
		_, terr = u.orderRepo.PaymentGetByCode(ctx, code)
		if terr == nil {
			code = utstring.RandomString(20)
			continue
		} else {
			if terr.GetType() != terror.ERROR_TYPE_DATA_NOT_FOUND {
				return
			}
		}
		terr = nil
		break
	}

	operator, terr := u.orderRepo.AuthGetFromContext(ctx)
	if terr != nil {
		return
	}

	paymentReq := models.Payment{
		Price:     totalPrice,
		Code:      code,
		ExpiredAt: time.Now().Add(time.Hour),
		UserID:    operator.ID,
	}

	var payment models.Payment
	payment, terr = u.orderRepo.PaymentCreate(ctx, paymentReq)
	if terr != nil {
		return
	}

	err := utstruct.InjectStructValue(payment, &paymentOrders)
	if err != nil {
		terr = terror.New(err)
		return
	}

	var orders []models.Order
	// create order
	for k, v := range reqMap {
		_, terr = u.orderRepo.WarehouseProductLock(ctx, models.WarehouseProduct{ProductID: k})
		if terr != nil {
			return
		}

		var total int64
		total, terr = u.orderRepo.WarehouseProductTotal(ctx, k)
		if terr != nil {
			return
		}

		if total < v {
			terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s is out of stock", products[k].Name))
			return
		}

		var order models.Order
		order, terr = u.orderRepo.OrderCreate(ctx, models.Order{
			PaymentID: int64(payment.ID),
			ProductID: k,
			Amount:    v,
		})
		if terr != nil {
			return
		}

		orders = append(orders, order)
	}

	paymentOrders.Orders = orders

	return
}

func (u OrderUsecase) OrderPayByCode(ctx *gin.Context, code string) (paymentOrder models.PaymentWithOrder, terr terror.ErrInterface) {
	helper.TxCreate(ctx, u.orderRepo.GetDB)
	defer func() {
		if terr != nil {
			helper.TxRollBack(ctx)
		} else {
			helper.TxCommit(ctx)
		}
	}()

	var payment models.Payment
	{
		payment, terr = u.orderRepo.PaymentGetByCode(ctx, code)
		if terr != nil {
			return
		}

		_, terr = u.orderRepo.PaymentLock(ctx, int64(payment.ID))
		if terr != nil {
			return
		}

		currentTime := time.Now()
		if currentTime.After(payment.ExpiredAt) {
			terr = terror.ErrInvalidRule("Payment has been expired")
			return
		}

		payment.Channel = "QRIS"
		payment.PaidOff = 1

		terr = u.orderRepo.OrderPaymentUpdate(ctx, payment)
		if terr != nil {
			return
		}
	}

	{
		var orders []models.Order
		orders, terr = u.orderRepo.OrdersGetByPaymentID(ctx, int64(payment.ID))
		if terr != nil {
			return
		}

		for _, v := range orders {
			_, terr = u.orderRepo.WarehouseProductLock(ctx, models.WarehouseProduct{ProductID: v.ProductID})
			if terr != nil {
				return
			}

			var total int64
			total, terr = u.orderRepo.WarehouseProductTotal(ctx, v.ProductID)
			if terr != nil {
				return
			}

			var pd models.Product
			pd, terr = u.orderRepo.ProductGetByID(ctx, v.ProductID)
			if terr != nil {
				return
			}

			if total < v.Amount {
				terr = terror.ErrInvalidRule(fmt.Sprintf("Product %s is out of stock", pd.Name))
				return
			}

			_, terr = u.orderRepo.WarehouseCustomerBuy(ctx, v)
			if terr != nil {
				return
			}
		}
	}

	return
}

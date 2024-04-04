package repository

import (
	"errors"

	"github.com/DeniesKresna/skyshi1gin/service/extensions/helper"
	"github.com/DeniesKresna/skyshi1gin/service/extensions/terror"
	"github.com/DeniesKresna/skyshi1gin/types/constants"
	"github.com/DeniesKresna/skyshi1gin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *OrderRepository) GetDB(ctx *gin.Context) (tx interface{}) {
	return r.db
}

func (r *OrderRepository) OrderList(ctx *gin.Context) (order []models.Order, terr terror.ErrInterface) {
	err := r.db.Find(&order).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrderGetByID(ctx *gin.Context, id int64) (order models.Order, terr terror.ErrInterface) {
	err := r.db.Take(&order, "id = ?", id).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrderCreate(ctx *gin.Context, req models.Order) (order models.Order, terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	err := tx.Create(&req).Error
	if err != nil {
		terr = terror.New(err)
	}
	order = req
	return
}

func (r *OrderRepository) PaymentCreate(ctx *gin.Context, req models.Payment) (payment models.Payment, terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	err := tx.Create(&req).Error
	if err != nil {
		terr = terror.New(err)
	}
	payment = req
	return
}

func (r *OrderRepository) PaymentGetByCode(ctx *gin.Context, code string) (payment models.Payment, terr terror.ErrInterface) {
	err := r.db.Take(&payment, "code = ?", code).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) PaymentLock(ctx *gin.Context, paymentID int64) (payment models.Payment, terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	err := tx.Clauses(clause.Locking{
		Strength: constants.TX_SHARE,
		Table:    clause.Table{Name: clause.CurrentTable},
	}).Where("id = ?", paymentID).First(&payment).Error

	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrdersGetByPaymentID(ctx *gin.Context, paymentID int64) (orders []models.Order, terr terror.ErrInterface) {
	err := r.db.Where("payment_id = ?", paymentID).Find(&orders).Error
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r *OrderRepository) OrderPaymentUpdate(ctx *gin.Context, req models.Payment) (terr terror.ErrInterface) {
	tx := helper.TxGet(ctx)
	if tx == nil {
		tx = r.db
	}

	err := tx.Updates(req).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

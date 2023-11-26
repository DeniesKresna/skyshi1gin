package repository

import (
	"fmt"

	"github.com/DeniesKresna/bengkelgin/service/extensions/terror"
	"github.com/DeniesKresna/bengkelgin/types/constants"
	"github.com/DeniesKresna/bengkelgin/types/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (r BengkelRepository) ExecutionSearch(ctx *gin.Context, exec models.ExecutionFilter, searchPayload models.DbSearchObject) (execs []models.ExecutionSearchResponse, totalData int64, terr terror.ErrInterface) {
	fmt.Printf("search.Payload: %+v\n", searchPayload)

	// raw
	queryDB := r.db.
		Joins("inner join (select * from users where deleted_at is null) as u on e.executer_id = u.id").
		Joins("inner join (select * from customers where deleted_at is null) as cu on e.customer_id = cu.id").
		Joins("inner join (select * from cars where deleted_at is null) as c on e.car_id = c.id").
		Where("e.deleted_at is null").Table("executions as e")

	// filter
	{
		if exec.CarPlat != "" {
			queryDB = queryDB.Where("c.plat like ?", exec.CarPlat)
		}
		if exec.CustomerName != "" {
			queryDB = queryDB.Where("cu.name like ?", exec.CustomerName)
		}
		if exec.CustomerEmail != "" {
			queryDB = queryDB.Where("cu.email like ?", exec.CustomerEmail)
		}
		if exec.ExecutedFrom != nil && exec.ExecutedTo != nil {
			queryDB = queryDB.Where("(e.executed_at between ? and ?)", *exec.ExecutedFrom, *exec.ExecutedTo)
		}
		if exec.ExecuterName != "" {
			queryDB = queryDB.Where("u.name like ?", exec.ExecuterName)
		}
		if exec.FinishFrom != nil && exec.FinishTo != nil {
			queryDB = queryDB.Where("(e.finish_at between ? and ?)", *exec.FinishFrom, *exec.FinishTo)
		}
		if len(exec.Paid) > 1 {
			queryDB = queryDB.Where("(e.paid between ? and ?)", exec.Paid[0], exec.Paid[1])
		}
		if len(exec.Price) > 1 {
			queryDB = queryDB.Where("(e.price between ? and ?)", exec.Price[0], exec.Price[1])
		}
		if exec.PaidFrom != nil && exec.PaidTo != nil {
			queryDB = queryDB.Where("(e.paid_at between ? and ?)", *exec.PaidFrom, *exec.PaidTo)
		}
		if exec.PaidOff != nil {
			paidOff := 0
			if *exec.PaidOff {
				paidOff = 1
			}
			queryDB = queryDB.Where("(e.paid_off = ?)", paidOff)
		}
	}

	queryDB = queryDB.Session(&gorm.Session{})

	if searchPayload.Mode == constants.DB_MODE_DATA || searchPayload.Mode == constants.DB_MODE_PAGE {
		offset := (searchPayload.Page - 1) * searchPayload.Limit
		queryDB = queryDB.Select("e.*", "u.name as user_name", "u.id as user_id", "c.plat", "cu.name as customer_name", "cu.phone as customer_phone", "cu.email as customer_email")
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		err := queryDB.Limit(int(searchPayload.Limit)).Offset(int(offset)).Find(&execs).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_COUNT || searchPayload.Mode == constants.DB_MODE_PAGE {
		err := queryDB.Limit(-1).Offset(-1).Count(&totalData).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	if searchPayload.Mode == constants.DB_MODE_DOWNLOAD {
		for _, v := range searchPayload.Order {
			queryDB = queryDB.Order(v)
		}
		err := queryDB.Select("e.*", "u.name as user_name", "u.id as user_id", "c.plat", "cu.name as customer_name", "cu.phone as customer_phone", "cu.email as customer_email").
			Find(&execs).Error
		if err != nil {
			terr = terror.New(err)
			return
		}
	}

	return
}

func (r BengkelRepository) ExecutionCreate(ctx *gin.Context, exec *models.Execution) (terr terror.ErrInterface) {
	err := r.db.Create(exec).Error
	if err != nil {
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) ExecutionUpdate(ctx *gin.Context, exec *models.Execution) (terr terror.ErrInterface) {
	err := r.db.Model(&exec).Updates(exec).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

func (r BengkelRepository) ExecutionGetByID(ctx *gin.Context, id int64) (exec models.ExecutionSearchResponse, terr terror.ErrInterface) {
	err := r.db.Raw(`
		select e.*, u.name as user_name, u.id as user_id, c.plat, cu.name as customer_name, cu.phone as customer_phone, cu.email as customer_email
		from executions as e
		inner join (select * from users where deleted_at is null) as u on e.executer_id = u.id
		inner join (select * from customers where deleted_at is null) as cu on e.customer_id = cu.id
		inner join (select * from cars where deleted_at is null) as c on e.car_id = c.id
		where e.deleted_at is null
	`).Where("e.id = ?", id).Model(&models.ExecutionSearchResponse{}).First(&exec).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			terr = terror.ErrNotFoundData(err.Error())
			return
		}
		terr = terror.New(err)
	}
	return
}

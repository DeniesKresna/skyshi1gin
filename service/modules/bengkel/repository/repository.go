package repository

import "gorm.io/gorm"

type BengkelRepository struct {
	db *gorm.DB
}

func BengkelCreateRepository(db *gorm.DB) BengkelRepository {
	bengkelRepository := BengkelRepository{
		db: db,
	}
	return bengkelRepository
}

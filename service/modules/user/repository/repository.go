package repository

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func UserCreateRepository(db *gorm.DB) UserRepository {
	userRepository := UserRepository{
		db: db,
	}
	return userRepository
}

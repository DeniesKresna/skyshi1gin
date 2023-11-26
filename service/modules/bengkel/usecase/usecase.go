package usecase

import (
	"github.com/DeniesKresna/bengkelgin/service/modules/bengkel/repository"
	userUsecase "github.com/DeniesKresna/bengkelgin/service/modules/user/usecase"
)

type BengkelUsecase struct {
	bengkelRepo repository.BengkelRepository
	userUsecase userUsecase.UserUsecase
}

func BengkelCreateUsecase(bengkelRepo repository.BengkelRepository, userUsecase userUsecase.UserUsecase) BengkelUsecase {
	bengkelUsecase := BengkelUsecase{
		bengkelRepo: bengkelRepo,
		userUsecase: userUsecase,
	}
	return bengkelUsecase
}

package di

import (
	"github.com/boon-neko/golang-grpc-backend-example/internal/infrastructure/database/repository"
	"github.com/boon-neko/golang-grpc-backend-example/internal/usecase"
)

type DIBox struct {
	UserUsecase usecase.UserUsecase
}

func NewDIBox() *DIBox {
	d := &DIBox{}
	d.TempDI()
	return d
}

func (d *DIBox) TempDI() {
	// Repository
	ur := repository.NewUser()
	// Usecase
	d.UserUsecase = usecase.NewUserInteractor(ur)

}

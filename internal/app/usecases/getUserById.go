package usecase

import (
	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/repository"
)

type UserUseCase struct {
    Repo repository.UserRepository
}

func (u *UserUseCase) GetUserByID(id int) (*entity.User, error) {
    return u.Repo.GetByID(id)
}

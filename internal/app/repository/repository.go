package repository

import "github.com/viniciusbls9/sofascore-api/internal/app/entity"

type UserRepository interface {
  GetByID(id int) (*entity.User, error)
}

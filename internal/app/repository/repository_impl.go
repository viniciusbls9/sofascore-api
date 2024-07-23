package repository

import (
	"database/sql"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/repository"
)

type UserRepositoryImpl struct {
    DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) repository.UserRepository {
    return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) GetByID(id int) (*entity.User, error) {
    user := &entity.User{}
    err := u.DB.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return nil, err
    }
    return user, nil
}

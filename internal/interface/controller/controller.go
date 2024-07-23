package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	usecase "github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

type UserController struct {
    UserUseCase *usecase.UserUseCase
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    user, err := uc.UserUseCase.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(user)
}

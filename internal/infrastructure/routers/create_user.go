package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerCreateUser(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerCreateUserUseCase(w, r)
}

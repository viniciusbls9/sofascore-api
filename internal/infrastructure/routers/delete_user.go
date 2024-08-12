package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerDeleteUserUseCase(w, r)
}

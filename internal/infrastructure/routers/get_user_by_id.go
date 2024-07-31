package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerGetUserByID(w, r)
}

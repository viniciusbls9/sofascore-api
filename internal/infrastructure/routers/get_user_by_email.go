package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerGetUserByEmail(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerGetUserByEmail(w, r)
}

package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerGetUsers(w, r)
}

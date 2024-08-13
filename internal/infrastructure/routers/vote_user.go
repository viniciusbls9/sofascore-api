package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/usecases"
)

func HandlerVoteUser(w http.ResponseWriter, r *http.Request) {
	usecases.HandlerVoteUser(w, r)
}

package routers

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
)

func HandlerHealthz(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, struct{}{})
}

package usecases

import (
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
)

func HandlerGetUserByEmail(w http.ResponseWriter, r *http.Request) {

	userEmail := r.URL.Query().Get("email")
	loggedInUserID := r.URL.Query().Get("logged_in_user_id")

	if loggedInUserID == "" {
		utils.RespondWithError(w, http.StatusInternalServerError, "loggedIn user not foud")
		return
	}

	user, err := GetUserByEmail(userEmail, loggedInUserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

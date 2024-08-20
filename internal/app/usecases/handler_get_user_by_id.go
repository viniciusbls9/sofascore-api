package usecases

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
)

func HandlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	loggedInUserID := r.URL.Query().Get("logged_in_user_id")

	if loggedInUserID == "" {
		utils.RespondWithError(w, http.StatusInternalServerError, "loggedIn user not foud")
		return
	}

	user, err := GetUserByID(userID, loggedInUserID)
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

package usecases

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerDeleteUserUseCase(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	userID := chi.URLParam(r, "userID")

	existingUser, err := GetUserByID(userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to check user existence: %v", err))
		return
	}

	if existingUser == nil {
		utils.RespondWithError(w, http.StatusConflict, "User not found")
		return
	}

	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	stmt, err := db.Query("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()

	utils.RespondWithJSON(w, http.StatusOK, user)
}

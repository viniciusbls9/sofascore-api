package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerCreateUserUseCase(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	existingUser, err := GetUserByEmail(user.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to check user existence: %v", err))
		return
	}

	if existingUser != nil {
		utils.RespondWithError(w, http.StatusConflict, "User with this email already exists")
		return
	}

	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users (id, name, email, fav_position, biography, image_url) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.New(), user.Name, user.Email, user.Fav_position, user.Biography, user.Image_url)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't execute statement: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, user)
}

package usecases

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerGetUserByID(w http.ResponseWriter, r *http.Request) {
	var result []entity.User
	userID := chi.URLParam(r, "userID")
	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't open database connection: %v", err))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email, velocity, fav_position, rating, biography, created_at FROM users WHERE id=$1", userID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't query DB: %v", err))
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Velocity, &user.Fav_position, &user.Rating, &user.Biography, &user.Created_at); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan rows DB: %v", err))
			return
		}
		result = append(result, user)
	}
	utils.RespondWithJSON(w, http.StatusOK, result)
}

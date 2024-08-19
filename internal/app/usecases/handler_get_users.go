package usecases

import (
	"fmt"
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	loggedInUserID := r.URL.Query().Get("logged_user_id")

	var users []entity.User

	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't query DB: %v", err))
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Velocity, &user.Fav_position, &user.Rating, &user.Biography, &user.Created_at, &user.Image_url); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan rows DB: %v", err))
			return
		}

		user.Has_voted, err = hasUserVoted(db, loggedInUserID, user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to check vote status: %v", err))
			return
		}

		users = append(users, user)
	}
	utils.RespondWithJSON(w, http.StatusOK, users)
}

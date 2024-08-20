package usecases

import (
	"fmt"
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
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
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Fav_position, &user.Biography, &user.Created_at, &user.Image_url); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't scan rows DB: %v", err))
			return
		}

		// Get the average user votes
		averageVotes, err := getAverageVotes(db, user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to calculate average votes: %v", err))
			return
		}

		user.AverageVotes = averageVotes

		users = append(users, user)
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

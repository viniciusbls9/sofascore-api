package usecases

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	var users []entity.User
	loggedInUserID := r.URL.Query().Get("logged_user_id")

	if loggedInUserID == "" {
		utils.RespondWithError(w, http.StatusInternalServerError, "loggedIn user not foud")
		return
	}

	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	err = validateIfLoggedInUserExists(db, loggedInUserID)

	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't query DB: %v", err))
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Fav_position, &user.Biography, &user.Image_url, &user.Age, &user.Height, &user.Preferred_foot, &user.Shirt_number, &user.Created_at); err != nil {
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

		// Obter a nota dada pelo usuário logado
		userVote, err := getUserVote(db, loggedInUserID, user.ID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get user vote: %v", err))
			return
		}
		user.CurrentUserVotes = userVote

		users = append(users, user)
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

func validateIfLoggedInUserExists(db *sql.DB, loggedInUserID string) error {
	var id string
	row := db.QueryRow("SELECT id FROM users WHERE id = $1", loggedInUserID)

	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("logged in user does not exist")
		}
		return fmt.Errorf("failed to validate logged in user: %v", err)
	}
	return nil
}

package usecases

import (
	"database/sql"
	"fmt"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func GetUserByID(userID string, loggedInUserID string) (*entity.User, error) {
	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("couldn't open database connection: %v", err)
	}
	defer db.Close()

	var user entity.User
	err = db.QueryRow("SELECT id, name, email, fav_position, biography, created_at, image_url, age, height, preferred_foot, shirt_number FROM users WHERE id=$1", userID).Scan(
		&user.ID, &user.Name, &user.Email, &user.Fav_position, &user.Biography, &user.Image_url, &user.Age, &user.Height, &user.Preferred_foot, &user.Shirt_number, &user.Created_at)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found: %v", err)
		}
		return nil, fmt.Errorf("couldn't query DB: %v", err)
	}

	if loggedInUserID != "" {
		vote, err := getUserVote(db, loggedInUserID, user.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get user vote: %v", err)
		}
		user.AverageVotes = vote
		user.CurrentUserVotes = vote
	}

	return &user, nil
}

package usecases

import (
	"fmt"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func GetUserByEmail(email string) (*entity.User, error) {
	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var user entity.User
	err = db.QueryRow("SELECT id, name, email, fav_position, biography, created_at, image_url, age, height, preferred_foot, shirt_number FROM users WHERE email=$1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Fav_position, &user.Biography, &user.Created_at, &user.Image_url, &user.Age, &user.Height, &user.Preferred_foot, &user.Shirt_number)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil // No user found
		}
		return nil, fmt.Errorf("failed to query user: %v", err)
	}

	// if loggedInUserID != "" {
	// 	vote, err := getUserVote(db, loggedInUserID, user.ID)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to get user vote: %v", err)
	// 	}
	// 	user.AverageVotes = vote
	// }

	return &user, nil
}

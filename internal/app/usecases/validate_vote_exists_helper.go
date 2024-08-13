package usecases

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func ValidateVoteExists(db *sql.DB, voterID, votedUserID uuid.UUID) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM votes WHERE voter_id = $1 AND voted_user_id = $2)", voterID, votedUserID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("couldn't check existing vote: %v", err)
	}

	return exists, nil
}

func RemoveVote(db *sql.DB, voterID, votedUserID uuid.UUID) error {
	_, err := db.Exec("DELETE FROM votes WHERE voter_id = $1 AND voted_user_id = $2", voterID, votedUserID)
	if err != nil {
		return fmt.Errorf("couldn't remove vote: %v", err)
	}

	_, err = db.Exec("UPDATE users SET rating = rating - 1 WHERE id = $1", votedUserID)
	if err != nil {
		return fmt.Errorf("couldn't update user rating after vote removal: %v", err)
	}

	return nil
}

package usecases

import "database/sql"

func hasUserVoted(db *sql.DB, voterID string, votedUserID string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM votes WHERE voter_id = $1 AND voted_user_id = $2)", voterID, votedUserID).Scan(&exists)
	return exists, err
}

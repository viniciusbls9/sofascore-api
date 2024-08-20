package usecases

import (
	"database/sql"
	"fmt"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
)

func getAverageVotes(db *sql.DB, userID string) (entity.AverageVotes, error) {
	var (
		totalShotVote     int
		totalMarkingVote  int
		totalQualityVote  int
		totalVelocityVote int
		totalVotes        int
	)

	err := db.QueryRow(`
		SELECT
			COALESCE(SUM(shot_vote), 0) AS total_shot_vote,
			COALESCE(SUM(marking_vote), 0) AS total_marking_vote,
			COALESCE(SUM(quality_vote), 0) AS total_quality_vote,
			COALESCE(SUM(velocity_vote), 0) AS total_velocity_vote,
			COUNT(*) AS total_votes
		FROM votes
		WHERE voted_user_id = $1
	`, userID).Scan(&totalShotVote, &totalMarkingVote, &totalQualityVote, &totalVelocityVote, &totalVotes)

	if err != nil {
		return entity.AverageVotes{}, fmt.Errorf("couldn't get average votes: %v", err)
	}

	var averageVotes entity.AverageVotes

	if totalVotes > 0 {
		averageVotes.ShotVote = float64(totalShotVote) / float64(totalVotes)
		averageVotes.MarkingVote = float64(totalMarkingVote) / float64(totalVotes)
		averageVotes.QualityVote = float64(totalQualityVote) / float64(totalVotes)
		averageVotes.VelocityVote = float64(totalVelocityVote) / float64(totalVotes)
		averageVotes.OverallAverage = (averageVotes.ShotVote + averageVotes.MarkingVote + averageVotes.QualityVote + averageVotes.VelocityVote) / 4
	}

	return averageVotes, nil
}

// getUserVote retorna a nota dada por um usuário específico para um usuário alvo
func getUserVote(db *sql.DB, voterID string, votedUserID string) (entity.AverageVotes, error) {
	var vote entity.AverageVotes

	err := db.QueryRow(`
		SELECT pass_vote, shot_vote, marking_vote, quality_vote, velocity_vote
		FROM votes
		WHERE voter_id = $1 AND voted_user_id = $2
	`, voterID, votedUserID).Scan(&vote.PassVote, &vote.ShotVote, &vote.MarkingVote, &vote.QualityVote, &vote.VelocityVote)

	if err != nil {
		if err == sql.ErrNoRows {
			return vote, nil
		}
		return vote, fmt.Errorf("couldn't get user vote: %v", err)
	}

	return vote, nil
}

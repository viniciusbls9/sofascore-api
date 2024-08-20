package usecases

import (
	"database/sql"
	"fmt"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
)

// Função para calcular a média dos votos e a média geral
func getAverageVotes(db *sql.DB, userID string) (entity.AverageVotes, error) {
	var averageVotes entity.AverageVotes

	rows, err := db.Query(`
		SELECT
			COALESCE(AVG(shot_vote), 0) AS average_shot,
			COALESCE(AVG(marking_vote), 0) AS average_marking,
			COALESCE(AVG(quality_vote), 0) AS average_quality,
			COALESCE(AVG(velocity_vote), 0) AS average_velocity
		FROM votes
		WHERE voted_user_id = $1
	`, userID)
	if err != nil {
		return averageVotes, fmt.Errorf("couldn't get average votes: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&averageVotes.ShotVote, &averageVotes.MarkingVote, &averageVotes.QualityVote, &averageVotes.VelocityVote)
		if err != nil {
			return averageVotes, fmt.Errorf("couldn't scan average votes: %v", err)
		}

		totalVotes := 0
		totalSum := 0.0

		if averageVotes.ShotVote > 0 {
			totalVotes++
			totalSum += averageVotes.ShotVote
		}
		if averageVotes.MarkingVote > 0 {
			totalVotes++
			totalSum += averageVotes.MarkingVote
		}
		if averageVotes.QualityVote > 0 {
			totalVotes++
			totalSum += averageVotes.QualityVote
		}
		if averageVotes.VelocityVote > 0 {
			totalVotes++
			totalSum += averageVotes.VelocityVote
		}

		if totalVotes > 0 {
			averageVotes.OverallAverage = totalSum / float64(totalVotes)
		}
	}

	return averageVotes, nil
}

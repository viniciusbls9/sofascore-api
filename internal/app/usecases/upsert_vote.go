package usecases

import (
	"database/sql"
	"fmt"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
)

func UpsertVote(db *sql.DB, voteRequest entity.VoteRequest) error {
	// Verifica se o voto já existe
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM votes WHERE voter_id = $1 AND voted_user_id = $2)", voteRequest.VoterID, voteRequest.VotedUserID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("couldn't check existing vote: %v", err)
	}

	if exists {
		// Atualiza o voto existente
		_, err = db.Exec("UPDATE votes SET pass_vote = $1, shot_vote = $2, marking_vote = $3, quality_vote = $4, velocity_vote = $5  WHERE voter_id = $6 AND voted_user_id = $7",
			voteRequest.PassVote, voteRequest.ShotVote, voteRequest.MarkingVote, voteRequest.QualityVote, voteRequest.VelocityVote, voteRequest.VoterID, voteRequest.VotedUserID)
		if err != nil {
			return fmt.Errorf("couldn't update vote: %v", err)
		}
	} else {
		// Insere um novo voto
		_, err = db.Exec("INSERT INTO votes (voter_id, voted_user_id, pass_vote, shot_vote, marking_vote, quality_vote, velocity_vote) VALUES ($1, $2, $3, $4, $5, $6, $7)",
			voteRequest.VoterID, voteRequest.VotedUserID, voteRequest.PassVote, voteRequest.ShotVote, voteRequest.MarkingVote, voteRequest.QualityVote, voteRequest.VelocityVote)
		if err != nil {
			return fmt.Errorf("couldn't insert vote: %v", err)
		}
	}

	return nil
}

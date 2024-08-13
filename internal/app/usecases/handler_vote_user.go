package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/viniciusbls9/sofascore-api/internal/app/entity"
	"github.com/viniciusbls9/sofascore-api/internal/app/utils"
	"github.com/viniciusbls9/sofascore-api/internal/infrastructure/db"
)

func HandlerVoteUser(w http.ResponseWriter, r *http.Request) {
	var voteRequest entity.VoteRequest

	err := json.NewDecoder(r.Body).Decode(&voteRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid JSON: %v", err))
		return
	}

	if voteRequest.VoterID == voteRequest.VotedUserID {
		utils.RespondWithError(w, http.StatusForbidden, "You cannot vote for yourself")
		return
	}

	db, err := db.HandlerOpenDatabaseConnection()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to connect to database: %v", err))
		return
	}
	defer db.Close()

	exists, err := ValidateVoteExists(db, voteRequest.VoterID, voteRequest.VotedUserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if exists {
		err := RemoveVote(db, voteRequest.VoterID, voteRequest.VotedUserID)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Vote removed successfully"})
		return
	}

	stmt, err := db.Prepare("INSERT INTO votes (voter_id, voted_user_id) VALUES ($1, $2)")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't prepare statement: %v", err))
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(voteRequest.VoterID, voteRequest.VotedUserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't execute statement: %v", err))
		return
	}

	_, err = db.Exec("UPDATE users SET rating = rating + 1 WHERE id = $1", voteRequest.VotedUserID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't update user rating: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Vote successful"})
}

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

	err = UpsertVote(db, voteRequest)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't upsert vote: %v", err))
		return
	}

	user, err := GetUserByID(voteRequest.VotedUserID.String())
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve user: %v", err))
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}

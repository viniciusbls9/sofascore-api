package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Velocity     int16     `json:"velocity"`
	Fav_position string    `json:"fav_position"`
	Rating       int16     `json:"rating"`
	Biography    string    `json:"biography"`
	Image_url    string    `json:"image_url"`
	Created_at   time.Time `json:"created_at"`
}

type VoteRequest struct {
	VoterID     uuid.UUID `json:"voter_id"`
	VotedUserID uuid.UUID `json:"voted_user_id"`
}

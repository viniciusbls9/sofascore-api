package entity

import (
	"time"

	"github.com/google/uuid"
)

type VoteRequest struct {
	VoterID      uuid.UUID `json:"voter_id"`
	VotedUserID  uuid.UUID `json:"voted_user_id"`
	PassVote     int16     `json:"pass_vote"`
	ShotVote     int16     `json:"shot_vote"`
	MarkingVote  int16     `json:"marking_vote"`
	QualityVote  int16     `json:"quality_vote"`
	VelocityVote int16     `json:"velocity_vote"`
}

type User struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Email            string       `json:"email"`
	Fav_position     string       `json:"fav_position"`
	Biography        string       `json:"biography"`
	Image_url        string       `json:"image_url"`
	Age              string       `json:"age"`
	Height           string       `json:"height"`
	Preferred_foot   string       `json:"preferred_foot"`
	Shirt_number     string       `json:"shirt_number"`
	Created_at       time.Time    `json:"created_at"`
	AverageVotes     AverageVotes `json:"average_votes"`
	CurrentUserVotes AverageVotes `json:"current_user_votes"`
}

type AverageVotes struct {
	PassVote       float64 `json:"pass_vote"`
	ShotVote       float64 `json:"shot_vote"`
	MarkingVote    float64 `json:"marking_vote"`
	QualityVote    float64 `json:"quality_vote"`
	VelocityVote   float64 `json:"velocity_vote"`
	OverallAverage float64 `json:"overall_average"`
}

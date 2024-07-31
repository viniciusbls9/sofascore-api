package entity

import "time"

type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Velocity     int16     `json:"velocity"`
	Fav_position string    `json:"fav_position"`
	Rating       int16     `json:"rating"`
	Biography    string    `json:"biography"`
	Created_at   time.Time `json:"created_at"`
}

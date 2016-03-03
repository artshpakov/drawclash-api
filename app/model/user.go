package model

import (
	"encoding/json"
	"github.com/artshpakov/drawclash/app"
	"time"
)

type User struct {
	ID        int
	Email     string
	CreatedAt time.Time
	Profile   *UserProfile
}

type UserProfile struct {
	Name   string
	Avatar string
}

func FindUser(id int64) (*User, error) {
	var user User
	_, err := app.GetDB().QueryOne(&user, `SELECT * FROM users WHERE id = ?`, id)
	return &user, err
}

func (u User) MarshalJSON() ([]byte, error) {
	v := struct {
		ID        int          `json:"id"`
		Email     string       `json:"email"`
		CreatedAt time.Time    `json:"created_at"`
		Profile   *UserProfile `json:"profile"`
	}{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		Profile:   u.Profile,
	}
	return json.Marshal(v)
}

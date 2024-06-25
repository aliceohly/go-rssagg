package main

import (
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatableAt time.Time `json:"updatable_at"`
	ApiKey      string    `json:"api_key"`
}

func dbUserToUser(user database.User) User {
	return User{
		ID:          user.ID,
		Name:        user.Name,
		CreatedAt:   user.CreatedAt,
		UpdatableAt: user.UpdatableAt,
		ApiKey:      user.ApiKey,
	}
}

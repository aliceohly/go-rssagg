// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Feed struct {
	ID            uuid.UUID
	Name          string
	Url           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        uuid.UUID
	LastFetchedAt time.Time
}

type FeedFollow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ApiKey    string
}

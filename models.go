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

type Feed struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Url         string    `json:"url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatableAt time.Time `json:"updatable_at"`
	UserID      uuid.UUID `json:"user_id"`
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

func dbFeedToFeed(feed database.Feed) Feed {
	return Feed{
		Name:        feed.Name,
		ID:          feed.ID,
		Url:         feed.Url,
		CreatedAt:   feed.CreatedAt,
		UpdatableAt: feed.UpdatableAt,
		UserID:      feed.UserID,
	}
}

func dbFeedsToFeeds(feeds []database.Feed) []Feed {
	var result []Feed
	for _, feed := range feeds {
		result = append(result, dbFeedToFeed(feed))
	}
	return result
}

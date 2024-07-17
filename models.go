package main

import (
	"database/sql"
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"` // or it will have error "unsupported Scan, storing driver.Value type <nil> into type *time.Time"
}

type FeedsFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func dbUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func dbFeedToFeed(feed database.Feed) Feed {
	return Feed{
		Name:          feed.Name,
		ID:            feed.ID,
		Url:           feed.Url,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		UserID:        feed.UserID,
		LastFetchedAt: nullTimeToTimePtr(feed.LastFetchedAt),
	}
}

func dbFeedsToFeeds(feeds []database.Feed) []Feed {
	var result []Feed
	for _, feed := range feeds {
		result = append(result, dbFeedToFeed(feed))
	}
	return result
}

func dbFeedFollowToFeedFollow(feedfollow database.FeedFollow) FeedsFollow {
	return FeedsFollow{
		ID:        feedfollow.ID,
		CreatedAt: feedfollow.CreatedAt,
		UpdatedAt: feedfollow.UpdatedAt,
		UserID:    feedfollow.UserID,
		FeedID:    feedfollow.FeedID,
	}
}

func dbFeedFollowsToFeedFollows(feedfollows []database.FeedFollow) []FeedsFollow {
	var result []FeedsFollow
	for _, feedfollow := range feedfollows {
		result = append(result, dbFeedFollowToFeedFollow(feedfollow))
	}
	return result
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

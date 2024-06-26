package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string
		Url  string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, "Could not decode parameters")
		return
	}

	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:          uuid.New(),
		Name:        params.Name,
		Url:         params.Url,
		CreatedAt:   time.Now().UTC(),
		UpdatableAt: time.Now().UTC(),
		UserID:      user.ID,
	})
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, dbFeedToFeed(feed))
}

func (cfg *apiConfig) handlerFeedsGet(w http.ResponseWriter, r *http.Request) {
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, dbFeedsToFeeds(feeds))
}

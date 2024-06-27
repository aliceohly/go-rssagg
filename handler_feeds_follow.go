package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId uuid.UUID
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, "Could not decode parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatableAt: time.Now().UTC(),
		UserID:      user.ID,
		FeedID:      params.FeedId,
	})
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, dbFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerFeedFollowGet(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := cfg.DB.GetFeedFollow(r.Context(), user.ID)
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, dbFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowString := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowString)
	if err != nil {
		responseWithErr(w, http.StatusBadRequest, "Invalid feedFollowId")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, "Could not delete feed follow")
		return
	}
	responseWithJSON(w, http.StatusOK, struct{}{})
}

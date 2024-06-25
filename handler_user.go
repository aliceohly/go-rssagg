package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aliceohly/go-rssagg/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUserCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, "Could not decode parameters")
		return
	}

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatableAt: time.Now(),
		Name:        params.Name,
	})
	if err != nil {
		responseWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, dbUserToUser(user))
}

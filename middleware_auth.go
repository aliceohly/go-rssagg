package main

import (
	"net/http"

	"github.com/aliceohly/go-rssagg/internal/auth"
	"github.com/aliceohly/go-rssagg/internal/database"
)

type authedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) middlewareAuth(authedHandler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responseWithErr(w, http.StatusUnauthorized, err.Error())
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithErr(w, http.StatusInternalServerError, err.Error())
			return
		}

		authedHandler(w, r, user)
	}
}

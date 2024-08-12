package main

import (
	"net/http"
	"fmt"
	"github.com/jikei25/rssagg/internal/auth"
	"github.com/jikei25/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig)middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		}
		handler(w, r, user)
	}
}
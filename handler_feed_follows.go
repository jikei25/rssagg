package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/jikei25/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameter{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow))
}

func (apiCfg *apiConfig)handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feed follows: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *apiConfig)handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feed_id_str := chi.URLParam(r, "feed_id")
	feed_id, err := uuid.Parse(feed_id_str)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't parse feed id: %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		FeedID: feed_id,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}
	
	respondWithJSON(w, 200, struct {
		Status string `json:"status"`
	}{
		Status : "Successful unfollow",
	})
}
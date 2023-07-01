package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/ombudhiraja/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameter struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameter{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, dbFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, mapper(feedFollows, dbFeedFollowToFeedFollow))

}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIdStr := chi.URLParam(r, "feedFollowId")

	feedFollowId, err := uuid.Parse(feedFollowIdStr)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalied FeedFollowId: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})

}

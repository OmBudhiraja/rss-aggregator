package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/ombudhiraja/rss-aggregator/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func dbUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserId:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func dbFeedFollowToFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeed.ID,
		UserId:    dbFeed.UserID,
		FeedId:    dbFeed.FeedID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Url         string    `json:"url"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func dbPostToPost(dbPost database.Post) Post {

	var description *string

	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}

	return Post{
		ID:          dbPost.ID,
		Title:       dbPost.Title,
		Description: description,
		Url:         dbPost.Url,
		PublishedAt: dbPost.PublishedAt,
		FeedID:      dbPost.FeedID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
	}
}

func mapper[S any, P any](slice []S, formatter func(S) P) []P {
	formatted := []P{}

	for _, value := range slice {
		formatted = append(formatted, formatter(value))
	}

	return formatted
}

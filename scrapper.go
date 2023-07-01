package main

import (
	"context"
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ombudhiraja/rss-aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error fetching feeds: ", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeRssFeed(db, wg, feed)
		}

		wg.Wait()
	}

}

func scrapeRssFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error Marking feed as fetched: ", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)

	if err != nil {
		log.Println("Error fetching feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {

		_, err := db.GetPostByURL(context.Background(), item.Link)

		if err == nil {
			continue
		}

		desc := sql.NullString{}

		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)

		if err != nil {
			log.Printf("Couldn't parse date %v with err: %v", item.PubDate, err)
			return
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			Title:       item.Title,
			Description: desc,
			Url:         item.Link,
			PublishedAt: pubAt,
			FeedID:      feed.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})

		if err != nil {
			log.Println("failed to create post")
		}
	}

	log.Printf("Feed `%s` collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

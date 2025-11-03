package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rowsedgy/gator/internal/database"
)

// func handlerAgg(s *state, cmd command) error {
// 	rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
// 	if err != nil {
// 		return fmt.Errorf("couldn't fetch feed: %w", err)
// 	}
// 	fmt.Printf("Feed: %+v\n", rssFeed)
// 	return nil
// }

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: agg <duration>")
	}

	time_between_reqs := cmd.Arguments[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing time_between_reqs: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching next feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:        nextFeed.ID,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return fmt.Errorf("error marking feed as fetched: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching rss feed from url: %w", err)
	}

	fmt.Println("Printing feed item titles:")
	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("* %s || Added on %v\n", item.Title, nextFeed.CreatedAt)
	}

	return nil
}

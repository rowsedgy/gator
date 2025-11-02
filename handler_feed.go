package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rowsedgy/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}
	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	f, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})

	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      s.cfg.CurrentUserName,
		Url:       feedUrl,
	})
	if err != nil {
		return fmt.Errorf("error following added feed: %w", err)
	}

	fmt.Println("Feed created succesfully:")
	printFeed(f)
	fmt.Println()
	fmt.Println("=====================================")

	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.ShowFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error fetching feeds from db: %w", err)
	}

	fmt.Println("Printing all feeds:")
	fmt.Println("----------------------------------")
	for _, feed := range feeds {
		fmt.Printf("* %s\n", feed.Name)
		fmt.Printf("* %s\n", feed.Url)
		fmt.Printf("* Added by %s\n", feed.FeedCreator)
		fmt.Println("----------------------------------")
	}
	return nil

}

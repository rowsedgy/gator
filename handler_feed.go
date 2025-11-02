package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rowsedgy/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedName := cmd.Arguments[0]
	feedUrl := cmd.Arguments[1]

	f, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    currentUser.ID,
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

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}

	activeUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedUrl := cmd.Arguments[0]

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      activeUser.Name,
		Url:       feedUrl,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	feedName, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("error getting feed name: %w", err)
	}

	fmt.Printf("Adding feed follow for user: %s\n", feedFollowRow.UserName)
	fmt.Printf("Feed followed: %s", feedName)

	return nil

}

func handlerFollowing(s *state, cmd command) error {
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	follows, err := s.db.GetFeedFollowsByUser(context.Background(), currentUser.Name)
	if err != nil {
		return fmt.Errorf("error getting follows for user %s: %w", currentUser.Name, err)
	}

	fmt.Printf("Feeds for current user (%s)\n", currentUser.Name)
	for _, follow := range follows {
		fmt.Printf("- '%s'\n", follow.FeedName)
	}

	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}

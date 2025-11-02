package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rowsedgy/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}

	feedUrl := cmd.Arguments[0]

	feedFollowRow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      user.Name,
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

func handlerFollowing(s *state, cmd command, user database.User) error {
	follows, err := s.db.GetFeedFollowsByUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("error getting follows for user %s: %w", user.Name, err)
	}

	fmt.Printf("Feeds for current user (%s)\n", user.Name)
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

func handlerFeedUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}

	feedUrl := cmd.Arguments[0]

	err := s.db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		UserID: user.ID,
		Url:    feedUrl,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	return nil
}

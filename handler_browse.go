package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rowsedgy/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int
	if len(cmd.Arguments) != 1 {
		limit = 2
	} else {
		var err error
		limit, err = strconv.Atoi(cmd.Arguments[0])
		if err != nil {
			return fmt.Errorf("error converting argument to int: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	})

	if err != nil {
		return fmt.Errorf("error getting posts for user: %w", err)
	}

	fmt.Printf("Printing posts for current user (%s). Limit set to %d:\n", user.Name, limit)
	for _, post := range posts {
		fmt.Printf("* Title: %s\n", post.Title)
		fmt.Printf("* Publish date: %s\n", post.PublishedAt)
		fmt.Printf("* URL: %s\n", post.Url)
		fmt.Printf("* Description: %s\n", post.Description)
		fmt.Println("=========================================")
	}

	return nil

}

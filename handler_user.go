package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rowsedgy/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Arguments[0]

	// Check if user exists
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("user %s doesn't exist\n%v", userName, err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("The user %s has been set.\n", userName)
	fmt.Printf("%+v", s.cfg)
	return nil

}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <user>", cmd.Name)
	}

	userName := cmd.Arguments[0]

	// Create user
	u, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	})

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User %s was created: %+v", userName, u)

	return nil
}

func handlerListUsers(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}

	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)

	}
	return nil
}

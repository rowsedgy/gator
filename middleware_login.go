package main

import "github.com/rowsedgy/gator/internal/database"

func middewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return nil
}

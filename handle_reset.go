package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.TruncateUsersTable(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("users table reset")
	return nil
}

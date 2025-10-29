package config

import (
	"errors"
	"fmt"
)

type state struct {
	State *Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Arguments) == 0 || len(cmd.Arguments) > 1 {
		return errors.New("no arguments provided")
	}

	userName := cmd.Arguments[0]

	err := s.State.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("The user %s has been set.\n", userName)
	return nil

}

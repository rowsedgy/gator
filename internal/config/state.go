package config

import (
	"errors"
	"fmt"
)

type State struct {
	State *Config
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Arguments) == 0 {
		return errors.New("no arguments provided")
	}

	userName := cmd.Arguments[0]

	err := s.State.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("The user %s has been set.\n", userName)
	fmt.Printf("%+v", s.State)
	return nil

}

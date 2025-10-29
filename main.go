package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rowsedgy/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v", cfg)

	state := config.State{
		State: &cfg,
	}

	commands := config.Commands{
		Commands: make(map[string]func(*config.State, config.Command) error),
	}

	userArgs := os.Args[1:]
	if len(userArgs) < 2 {
		log.Fatalf("at least one argument required")
	}

	command := config.Command{
		Name:      userArgs[1],
		Arguments: userArgs[2:],
	}

	commands.Register(command.Name, config.HandlerLogin)

}

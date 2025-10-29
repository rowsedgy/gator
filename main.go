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
	fmt.Printf("Read config: %+v\n", cfg)

	state := config.State{
		State: &cfg,
	}

	commands := config.Commands{
		Commands: make(map[string]func(*config.State, config.Command) error),
	}

	if len(os.Args) < 2 {
		log.Fatalf("at least one argument required")
	}

	command := config.Command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	commands.Register(command.Name, config.HandlerLogin)

	err = commands.Run(&state, command)
	if err != nil {
		log.Fatalf("error running command %v", err)
	}

}

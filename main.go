package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rowsedgy/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)

	programState := state{
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.Register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("at least one argument required")
	}

	command := command{
		Name:      os.Args[1],
		Arguments: os.Args[2:],
	}

	err = commands.Run(&programState, command)
	if err != nil {
		log.Fatal(err)
	}

}

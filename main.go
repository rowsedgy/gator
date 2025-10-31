package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rowsedgy/gator/internal/config"
	"github.com/rowsedgy/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error opening db %v", err)
	}

	dbQueries := database.New(db)

	programState := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}

	commands.Register("login", handlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", handlerReset)
	commands.Register("users", handlerListUsers)

	if len(os.Args) < 2 {
		log.Fatalf("at least one argument required")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.Run(&programState, command{Name: cmdName, Arguments: cmdArgs})
	if err != nil {
		fmt.Println("error: %v", err)
		os.Exit(1)
	}

}

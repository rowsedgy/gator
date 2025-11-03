package main

import (
	"context"
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
	// fmt.Printf("Read config: %+v\n", cfg)

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
	commands.Register("agg", handlerAgg)
	commands.Register("addfeed", middewareLoggedIn(handlerAddFeed))
	commands.Register("feeds", handlerFeeds)
	commands.Register("follow", middewareLoggedIn(handlerFollow))
	commands.Register("following", middewareLoggedIn(handlerFollowing))
	commands.Register("unfollow", middewareLoggedIn(handlerFeedUnfollow))
	commands.Register("browse", middewareLoggedIn(handlerBrowse))

	if len(os.Args) < 2 {
		log.Fatalf("at least one argument required")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.Run(&programState, command{Name: cmdName, Arguments: cmdArgs})
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

}

func middewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("user not logged-in: %w", err)
		}
		return handler(s, cmd, user)
	}
}

// func middlewareTicker(handler func(s *state, cmd command, time_between_req string) func(*state, command) error) {
// 	return func(s *state, cmd command) error {

// 	}
// }

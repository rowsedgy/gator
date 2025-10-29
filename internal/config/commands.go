package config

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	Commands map[string]func(*state, command) error
}

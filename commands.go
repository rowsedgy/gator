package main

type command struct {
	Name      string
	Arguments []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) Run(s *state, cmd command) error {
	com, ok := c.commands[cmd.Name]
	if ok {
		err := com(s, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *commands) Register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

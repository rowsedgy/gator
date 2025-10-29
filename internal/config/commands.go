package config

type Command struct {
	Name      string
	Arguments []string
}

type Commands struct {
	Commands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	com, ok := c.Commands[cmd.Name]
	if ok {
		err := com(s, cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	_, ok := c.Commands[name]
	if !ok {
		c.Commands[name] = f
	}
}

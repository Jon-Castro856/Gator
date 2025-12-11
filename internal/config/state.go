package config

import (
	"errors"

	"github.com/Jon-Castro856/Gator/internal/database"
)

type State struct {
	Config *Config
	DB     *database.Queries
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	RegisteredCommands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}

func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.RegisteredCommands[name] = f
}

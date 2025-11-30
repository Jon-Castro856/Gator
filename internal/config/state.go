package config

type State struct {
	Config *Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	Cmd map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {

}

func (c *Commands) register(name string, f func(*State, Command) error) {
	return
}

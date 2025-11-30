package main

import (
	"fmt"
	"os"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func main() {
	jsonConfig := config.Read()
	state := config.State{
		Config: &jsonConfig,
	}
	commandArgs := config.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	cmds := config.Commands{}

	cmds.Cmd["login"] = handlerLogin

	if os.Args < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}

	action := commandArgs.Name

	cmds.Run(cmds.Cmd[action](&state, commandArgs))

	return
}

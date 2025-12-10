package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/Jon-Castro856/Gator/internal/config"
)

func main() {
	jsonConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &config.State{
		Config: &jsonConfig,
	}

	cmds := config.Commands{
		RegisteredCommands: make(map[string]func(*config.State, config.Command) error),
	}
	cmds.Register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, config.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}

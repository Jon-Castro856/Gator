package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/Jon-Castro856/Gator/internal/config"
	"github.com/Jon-Castro856/Gator/internal/database"
)

func main() {
	const dbURL = "postgres://postgres:sql@localhost:5432/gator"
	jsonConfig, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("error establishing database connection: %v", err)
	}

	programState := &config.State{
		Config: &jsonConfig,
		DB:     database.New(db),
	}

	cmds := config.Commands{
		RegisteredCommands: make(map[string]func(*config.State, config.Command) error),
	}
	cmds.Register("login", handlerLogin)
	cmds.Register("register", handlerRegister)

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

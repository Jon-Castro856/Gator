package main

import (
	"github.com/Jon-Castro856/Gator/internal/config"
)

func main() {
	jsonConfig := config.Read()
	state := config.State{
		Config: &jsonConfig,
	}

}

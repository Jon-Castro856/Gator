package main

import (
	"github.com/Jon-Castro856/Gator/internal/config"
)

func main() {
	jsonConfig := config.Read()
	jsonConfig.SetUser("Jon")
	jsonConfig = config.Read()
}

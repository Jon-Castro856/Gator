package main

import "github.com/Jon-Castro856/Gator/internal/config"

func initCliCmds(cmd config.Commands) {
	cmd.Register("login", handlerLogin)
	cmd.Register("register", handlerRegister)
	cmd.Register("reset", handlerReset)
	cmd.Register("users", handlerGetUsers)
	cmd.Register("agg", handlerAgg)
	cmd.Register("addfeed", handlerAddFeed)
	cmd.Register("feeds", handlerFeeds)
}

package main

import "github.com/Jon-Castro856/Gator/internal/config"

func initCliCmds(cmd config.Commands) {
	cmd.Register("login", handlerLogin)
	cmd.Register("register", handlerRegister)
	cmd.Register("reset", handlerReset)
	cmd.Register("users", handlerGetUsers)
	cmd.Register("agg", handlerAgg)
	cmd.Register("addfeed", middlewareLoggedIn(handlerAddFeed))
	cmd.Register("feeds", handlerFeeds)
	cmd.Register("follow", middlewareLoggedIn(handlerFollow))
	cmd.Register("following", middlewareLoggedIn(handlerFollowing))
	cmd.Register("unfollow", middlewareLoggedIn(handlerUnFollow))
	cmd.Register("browse", handlerBrowse)
}

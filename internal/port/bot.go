package port

import (
	"github.com/NoBypass/fds/internal/common/env"
	"github.com/NoBypass/fds/internal/port/interaction"
	"github.com/NoBypass/fds/pkg/discord"
	"github.com/NoBypass/mincache"
	"github.com/labstack/gommon/log"
)

func RunBot(cfg *env.Env, cache *mincache.Cache) {
	session, err := discord.NewSession(cfg.DiscordToken, cache)
	if err != nil {
		log.Fatalf("error while creating discord session: %s", err)
	}

	err = session.RegisterCommands(
		&interaction.CmdPing{})
	if err != nil {
		log.Fatalf("error while registering commands: %s", err)
	}

	session.HandleEvents()
}

package discord

import (
	"fmt"
	"github.com/NoBypass/mincache"
	"github.com/bwmarrin/discordgo"
)

type Session struct {
	dcgo  *discordgo.Session
	cmds  map[string]Command
	cache *mincache.Cache
}

func NewSession(token string, cache *mincache.Cache) (*Session, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("problem while creating discordgo instance: %s", err)
	}
	err = s.Open()
	if err != nil {
		return nil, fmt.Errorf("error while opening discord session: %s", err)
	}

	return &Session{
		dcgo:  s,
		cmds:  make(map[string]Command),
		cache: cache,
	}, nil
}

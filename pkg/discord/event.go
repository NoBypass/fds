package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"time"
)

func (s *Session) HandleEvents() {
	s.dcgo.AddHandler(func(session *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if cmd, ok := s.cmds[i.ApplicationCommandData().Name]; ok {
				ctx, err := cmd.Run(s.parseCommandInteraction(i), context.Background())
				if err != nil {
					return // TODO implement error handler
				}

				s.cache.Set(i.Member.User.ID, ctx, time.Minute)
			}
		}
	})
}

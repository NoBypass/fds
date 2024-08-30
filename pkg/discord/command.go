package discord

import (
	"context"
	"github.com/bwmarrin/discordgo"
)

type Command interface {
	Run(i *CommandInteraction, ctx context.Context) (context.Context, error)
	Get() (string, *discordgo.ApplicationCommand)
}

func (s *Session) RegisterCommands(cmds ...Command) error {
	for _, cmd := range cmds {
		id, appCmd := cmd.Get()
		_, err := s.dcgo.ApplicationCommandCreate(s.dcgo.State.User.ID, "", appCmd)
		if err != nil {
			return err
		}
		s.cmds[id] = cmd
	}

	return nil
}

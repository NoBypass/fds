package lifecycle

import "github.com/bwmarrin/discordgo"

type Bot struct {
	Token          *string
	RemoveCommands *bool
	Session        *discordgo.Session
}

package cmds

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var Ping = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Get the bot's ping",
	Version:     "v1.1.0",
}

func PingHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	latency := s.HeartbeatLatency().Milliseconds()

	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Your ping to the bot (EU) is %vms", latency),
		},
	})
}

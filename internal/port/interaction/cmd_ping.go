package interaction

import (
	"context"
	"fmt"
	"github.com/NoBypass/fds/pkg/discord"
	"github.com/bwmarrin/discordgo"
)

type CmdPing struct {
}

func (c *CmdPing) Get() (string, *discordgo.ApplicationCommand) {
	return "ping", &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Ping the bot",
		Version:     "v1.2.1",
	}
}

func (c *CmdPing) Run(i *discord.CommandInteraction, ctx context.Context) (context.Context, error) {
	latency := i.Session().HeartbeatLatency().Milliseconds()
	return ctx, i.Respond(&discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("The bots latency is %vms", latency),
		},
	})
}

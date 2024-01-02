package cmds

import (
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/bwmarrin/discordgo"
)

var two = 2.0

var teamsPrinter = func(s *discordgo.Session, i *discordgo.InteractionCreate, teams [][]string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title: "Teams",
					Color: consts.EmbedColor,
					Fields: func() []*discordgo.MessageEmbedField {
						var fields []*discordgo.MessageEmbedField
						for i, team := range teams {
							var val string
							for _, player := range team {
								val += fmt.Sprintf("` %v `\n", player)
							}
							fields = append(fields, &discordgo.MessageEmbedField{
								Inline: true,
								Name:   fmt.Sprintf("Team %v", i+1),
								Value:  val,
							})
						}
						return fields
					}(),
				},
			},
		},
	})
}

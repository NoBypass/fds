package cmds

import (
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

var adminPerms = int64(discordgo.PermissionAdministrator)

var Admin = &discordgo.ApplicationCommand{
	Name:                     "admin",
	Description:              "Admin utilities",
	Version:                  "v1.0.0",
	DefaultMemberPermissions: &adminPerms,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "embed",
			Description: "Write an embed to the channel",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    false,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "verify",
					Value: "verify",
				},
				{
					Name:  "test",
					Value: "test",
				},
			},
		},
	},
}

func AdminHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	om := utils.OptionMap(i.ApplicationCommandData().Options)
	embed := om["embed"].(string)

	var res *discordgo.MessageSend

	switch embed {
	case "verify":
		res = &discordgo.MessageSend{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       "Verify",
					Color:       consts.EmbedColor,
					Description: "Verify your Discord account by linking it to Hypixel.",
				},
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							CustomID: "verify",
							Style:    discordgo.SuccessButton,
							Label:    "Verify",
							Emoji: discordgo.ComponentEmoji{
								Name: "🔗",
							},
						},
					},
				},
			},
		}
	case "test":
		res = &discordgo.MessageSend{
			Content: "Test",
			Embed: &discordgo.MessageEmbed{
				Title: "Test",
			},
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							CustomID: "test",
							Style:    discordgo.PrimaryButton,
							Label:    "Test",
						},
					},
				},
			},
		}
	default:
		res = &discordgo.MessageSend{
			Content: "",
		}
	}

	_, err := s.ChannelMessageSendComplex(i.ChannelID, res)
	if err != nil {
		return err
	}
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Message was sent to channel",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

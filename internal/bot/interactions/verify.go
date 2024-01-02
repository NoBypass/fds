package interactions

import "github.com/bwmarrin/discordgo"

func VerifyHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: "verify_modal_" + i.Interaction.Member.User.ID,
			Title:    "Verify " + i.Interaction.Member.User.Username,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.TextInput{
							CustomID:    "current_time",
							Label:       "What time is your Minecraft name?",
							Style:       discordgo.TextInputShort,
							Placeholder: "Your Minecraft name",
							Required:    true,
							MaxLength:   16,
							MinLength:   1,
						},
					},
				},
			},
		},
	})
}

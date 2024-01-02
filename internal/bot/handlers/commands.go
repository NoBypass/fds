package handlers

import (
	"fmt"
	"github.com/NoBypass/fds/internal/bot/cmds"
	"github.com/bwmarrin/discordgo"
	"log"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
	"ping":    cmds.PingHandler,
	"teams":   cmds.TeamsHandler,
	"vcteams": cmds.VCTeamsHandler,
	"admin":   cmds.AdminHandler,
}

var commands = []*discordgo.ApplicationCommand{
	cmds.Ping,
	cmds.Teams,
	cmds.VCTeams,
	cmds.Admin,
}

func RegisterCommands(s *discordgo.Session) {
	for _, c := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", c)
		if err != nil {
			log.Fatalf("Cannot register command %v: %v", c.Name, err)
		}
		log.Printf("Registered command %v", c.Name)
	}
}

func handleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		err := h(s, i)
		if err != nil {
			panic(err)
		}
	}
}

func respondErr(s *discordgo.Session, i *discordgo.InteractionCreate, err error) {
	e := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: fmt.Sprintf("Oops, something went wrong: %v\n\nIf this keeps happening, please contact staff as this is likely an easy fix", err),
		},
	})
	if e != nil {
		log.Printf("Cannot send error message: %v", e)
	}
}

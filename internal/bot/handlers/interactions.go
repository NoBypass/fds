package handlers

import (
	"fmt"
	"github.com/NoBypass/fds/internal/bot/interactions"
	"github.com/bwmarrin/discordgo"
	"log"
)

var interactionHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) error{
	"verify": interactions.VerifyHandler,
}

func Interactions(s *discordgo.Session, i *discordgo.InteractionCreate) {
	defer func() {
		if r := recover(); r != nil {
			r = fmt.Errorf("(recovered) panic: %v", r)
			respondErr(s, i, r.(error))
			log.Print(r)
		}
	}()

	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		handleCommand(s, i)
	case discordgo.InteractionMessageComponent:
		err := interactionHandlers[i.MessageComponentData().CustomID](s, i)
		if err != nil {
			panic(err)
		}
	default:
		log.Printf("Unknown interaction type: %v", i.Type)
		respondErr(s, i, fmt.Errorf("unknown interaction type: %v", i.Type))
	}
}

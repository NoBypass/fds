package main

import (
	"flag"
	"fmt"
	"github.com/NoBypass/fds/internal/bot/handlers"
	"github.com/NoBypass/fds/internal/bot/lifecycle"
	"github.com/NoBypass/fds/internal/pkg/conf"
	"github.com/NoBypass/fds/internal/pkg/consts"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
)

var (
	s              *discordgo.Session
	Config         = conf.ReadConfig()
	BotToken       = flag.String("token", Config.Authentication.Bot.Token, "Bot access token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
	b              = &lifecycle.Bot{
		Token:          BotToken,
		RemoveCommands: RemoveCommands,
	}
)

func init() { flag.Parse() }

func init() {
	var err error
	s, err = discordgo.New("Bot " + *BotToken)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
	b.Session = s
	log.Println("Session created")

	err = s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	log.Println("Session opened")
}

const VERSION = "v3.0.0"

func main() {
	defer s.Close()

	fmt.Println(`
   _______  ____   ___       __
  / __/ _ \/ __/  / _ )___  / /_
 / _// // /\ \   / _  / _ \/ __/
/_/ /____/___/  /____/\___/\__/   ` + consts.Purple.Sprint(VERSION) + `
The FDS Discord bot written in    ` + consts.WhiteOnCyan.Sprint(" GO ") + `
________________________________________________
`)

	handlers.RegisterCommands(s)

	s.AddHandler(handlers.Ready)
	s.AddHandler(handlers.Interactions)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	b.Shutdown()
}

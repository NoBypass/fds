package cmds

import (
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

var Teams = &discordgo.ApplicationCommand{
	Name:        "teams",
	Description: "Generate random teams",
	Version:     "v1.0.1",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "players",
			Description: "List of players seperated by a space",
			Type:        discordgo.ApplicationCommandOptionString,
			Required:    true,
		},
		{
			Name:        "teams",
			Description: "Number of teams (default: 2)",
			Type:        discordgo.ApplicationCommandOptionInteger,
			MinValue:    &two,
			Required:    false,
		},
		{
			Name:        "members",
			Description: "Number of members per team (takes priority over teams)",
			Type:        discordgo.ApplicationCommandOptionInteger,
			MinValue:    &two,
			Required:    false,
		},
	},
}

func TeamsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	om := utils.OptionMap(i.ApplicationCommandData().Options)
	playersStr := om["players"].(string)
	teamAmount, tOk := om["teams"].(float64)
	memberAmount, mOk := om["members"].(float64)
	players := strings.Split(playersStr, " ")
	playerAmount := len(players)

	if tOk && mOk {
		return fmt.Errorf("cannot define both memberAmount and teamAmount")
	} else if !tOk && !mOk {
		teamAmount = 2
	}

	var teams [][]string
	if memberAmount != 0 {
		teams = make([][]string, playerAmount/int(memberAmount))
	} else {
		teams = make([][]string, int(teamAmount))
	}

	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})
	for i, player := range players {
		teams[i%len(teams)] = append(teams[i%len(teams)], player)
	}

	return teamsPrinter(s, i, teams)
}

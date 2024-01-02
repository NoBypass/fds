package cmds

import (
	"fmt"
	"github.com/NoBypass/fds/internal/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

var VCTeams = &discordgo.ApplicationCommand{
	Name:        "vcteams",
	Description: "Generate random teams from the members in your voice channel",
	Version:     "v1.0.0",
	Options: []*discordgo.ApplicationCommandOption{
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

func VCTeamsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	om := utils.OptionMap(i.ApplicationCommandData().Options)
	teamAmount, tOk := om["teams"].(float64)
	memberAmount, mOk := om["members"].(float64)

	if tOk && mOk {
		return fmt.Errorf("cannot define both memberAmount and teamAmount")
	} else if !tOk && !mOk {
		teamAmount = 2
	}

	userID := i.Member.User.ID
	guildID := i.GuildID
	voiceState, err := s.State.VoiceState(guildID, userID)
	if err != nil {
		return err
	}
	voiceChannelID := voiceState.ChannelID
	guild, err := s.State.Guild(guildID)
	if err != nil {
		return err
	}

	var members []string
	for _, vs := range guild.VoiceStates {
		if vs.ChannelID == voiceChannelID {
			member, err := s.State.Member(guildID, vs.UserID)
			if err != nil {
				return err
			}
			members = append(members, member.Nick)
		}
	}

	var teams [][]string
	if memberAmount != 0 {
		teams = make([][]string, len(members)/int(memberAmount))
	} else {
		teams = make([][]string, int(teamAmount))
	}

	rand.Shuffle(len(members), func(i, j int) {
		members[i], members[j] = members[j], members[i]
	})
	for i, player := range members {
		teams[i%len(teams)] = append(teams[i%len(teams)], player)
	}

	return teamsPrinter(s, i, teams)
}

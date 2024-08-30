package discord

import "github.com/bwmarrin/discordgo"

type CommandInteraction struct {
	ID        string
	GuildID   string
	ChannelID string
	InDM      bool
	Locale    discordgo.Locale
	Member    *discordgo.Member
	User      *discordgo.User
	Args      map[string]*CommandOption
	Context   discordgo.ApplicationCommandType
	TargetID  string

	session     *discordgo.Session
	interaction *discordgo.Interaction
}

type CommandOption struct {
	Type    discordgo.ApplicationCommandOptionType
	Options map[string]*CommandOption
	Value   any
}

func (ci *CommandInteraction) Respond(content *discordgo.InteractionResponse) error {
	return ci.session.InteractionRespond(ci.interaction, content)
}

func (ci *CommandInteraction) Session() *discordgo.Session {
	return ci.session
}

func (s *Session) parseCommandInteraction(i *discordgo.InteractionCreate) *CommandInteraction {
	data := i.ApplicationCommandData()

	return &CommandInteraction{
		ID:        i.ID,
		GuildID:   i.GuildID,
		ChannelID: i.ChannelID,
		InDM:      i.Member == nil,
		Locale:    i.Locale,
		Member:    i.Member,
		User:      i.User,
		Args:      parseOptions(data.Options),
		Context:   data.CommandType,
		TargetID:  data.TargetID,

		session:     s.dcgo,
		interaction: i.Interaction,
	}
}

func parseOptions(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*CommandOption {
	opts := make(map[string]*CommandOption)
	for _, opt := range options {
		opts[opt.Name] = &CommandOption{
			Options: parseOptions(opt.Options),
			Value:   opt.Value,
			Type:    opt.Type,
		}
	}
	return opts
}

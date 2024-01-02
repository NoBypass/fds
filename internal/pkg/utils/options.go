package utils

import "github.com/bwmarrin/discordgo"

func OptionMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]interface{} {
	optionMap := make(map[string]interface{})
	for _, v := range options {
		optionMap[v.Name] = v.Value
	}
	return optionMap
}

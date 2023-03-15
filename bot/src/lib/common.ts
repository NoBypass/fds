import { CommandInteraction, Guild, TextChannel } from 'discord.js'

export const getTextChannel = (interaction: CommandInteraction, id?: string) => {
    return (interaction.guild as Guild).channels.resolve(id? id: interaction.channelId) as TextChannel
}
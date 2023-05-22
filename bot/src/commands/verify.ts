import {ChannelType, Guild, PermissionFlagsBits, SlashCommandBuilder, TextChannel} from 'discord.js'
import {SlashCommand} from '../types'

const ClearCommand: SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('verify')
        .setDescription('Link your discord with you minecraft account')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages)
        .addStringOption(option => {
            return option
                .setName('ign')
                .setDescription('Your Minecraft in game name')
                .setRequired(true)
            // option.setAutocomplete()
        }),
    execute: interaction => {
        const ign = interaction.options.get('ign')
    },
    cooldown: 10
}

export default ClearCommand
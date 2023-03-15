import { ChannelType, Guild, PermissionFlagsBits, SlashCommandBuilder, TextChannel } from 'discord.js'
import { SlashCommand } from '../types'
import Embed from '../lib/Embed'

const ClearCommand : SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('verify')
        .setDescription('Link your discord with you minecraft account')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {

    },
    cooldown: 10
}

export default ClearCommand
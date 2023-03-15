import { ChannelType, Guild, PermissionFlagsBits, SlashCommandBuilder, TextChannel } from 'discord.js'
import { SlashCommand } from '../types'
import Embed from '../lib/Embed'

const ClearCommand : SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('daily')
        .setDescription('Claim your daily exp through this command')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const embed = new Embed({ interaction })
        const xpToGive = Math.round(Math.random() * 500)

        embed.title(`${interaction.user.username} ${xpToGive < 400 && xpToGive > 100? 'claimed their daily reward': `got ${xpToGive > 450 || xpToGive < 50? 'very': ''} ${xpToGive < 100? 'un': ''}lucky`}`)
        embed.description(`Received **${xpToGive}**xp`)

        // TODO streaks
        // TODO check if level changed
        // TODO link with database for guild xp and just for giving the xp
    },
    cooldown: 10
}

export default ClearCommand
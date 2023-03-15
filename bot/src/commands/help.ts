import { ChannelType, PermissionFlagsBits, SlashCommandBuilder } from 'discord.js'
import { SlashCommand } from '../types'
import { join } from "path"
import { readdirSync } from "fs"
import Embed from "../lib/Embed"

const ClearCommand : SlashCommand = {
    command: new SlashCommandBuilder()
        .setName('help')
        .setDescription('View all the commands and some additional info')
        .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const commands : SlashCommandBuilder[] = []
        let commandsDir = join(__dirname,"../commands")
        readdirSync(commandsDir).forEach(file => {
            if (!file.endsWith(".ts")) return
            let command : SlashCommand = require(`${commandsDir}/${file}`).default
            commands.push(command.command)
        })

        const embed = new Embed()
        embed.title('Help Menu')
        embed.description(`Here's a list of all slash commands provided by the <@${interaction.client.user.id}>`)
        commands.forEach((command) => {
            embed.field({
                name: `/${command.name}`,
                value: command.description,
                inline: true
            })
        })
        interaction.reply({ embeds: [embed.get() as any], ephemeral: true })
    },
    cooldown: 10
}

export default ClearCommand
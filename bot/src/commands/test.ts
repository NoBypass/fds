import {SlashCommandBuilder, ChannelType, PermissionFlagsBits} from "discord.js"
import { SlashCommand } from "../types"
import Embed from "../lib/Embed"

const command : SlashCommand = {
    command: new SlashCommandBuilder()
    .setName('test')
    .setDescription('Shows the bot\'s ping and tests slash commands')
    .setDefaultMemberPermissions(PermissionFlagsBits.SendMessages),

    execute: interaction => {
        const embed = new Embed({ interaction })
        embed.description(`Test successful, Ping: ${interaction.client.ws.ping}`)
        return embed.send()
    },
    cooldown: 10
}

export default command
import { Client, Routes, SlashCommandBuilder } from "discord.js"
import { REST } from "@discordjs/rest"
import { readdirSync } from "fs"
import { join } from "path"
import { SlashCommand } from "../types"
import log from "../lib/log"

module.exports = (client : Client) => {
    const commands : SlashCommandBuilder[] = []
    let commandsDir = join(__dirname,"../commands")

    readdirSync(commandsDir).forEach(file => {
        if (!file.endsWith(".ts")) return
        let command : SlashCommand = require(`${commandsDir}/${file}`).default
        commands.push(command.command)
        client.slashCommands.set(command.command.name, command)
    })
    const rest = new REST({version: "10"}).setToken(process.env.TOKEN)

    log('', false)
    log('&c:blue;&s:underscore;Commands:', false)
    commands.forEach((command) => {
        log(`&c:yellow;Command '/${command.name}' registered`)
    })
    rest.put(Routes.applicationCommands(process.env.CLIENT_ID), {
        body: commands.map(command => command.toJSON())
    })
    .then(() => {
        log(`&c:green;Successfully loaded &bg:green;&c:black; ${commands.length} &s:reset;&c:green; command${commands.length == 1? '': 's'}`)
    }).catch(e => {
        log('&c:red;Encountered error during registration')
    })
}
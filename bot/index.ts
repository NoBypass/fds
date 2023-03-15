import { Client, Collection, GatewayIntentBits } from "discord.js"
import { Command, SlashCommand } from './src/types'
import { config } from 'dotenv'
import { readdirSync } from 'fs'
import { join } from 'path'
import log from "./src/lib/log"
const { Guilds, MessageContent, GuildMessages, GuildMembers } = GatewayIntentBits
export const client = new Client({ intents:[Guilds, MessageContent, GuildMessages, GuildMembers] })
config()

log('', false)
log('&bg:blue;&c:black; FDS Bot v2 is starting ', false)
client.slashCommands = new Collection<string, SlashCommand>()
client.commands = new Collection<string, Command>()
client.cooldowns = new Collection<string, number>()

const handlersDir = join(__dirname, "./src/handlers")
readdirSync(handlersDir).forEach(handler => {
    require(`${handlersDir}/${handler}`)(client)
})

client.login(process.env.TOKEN)
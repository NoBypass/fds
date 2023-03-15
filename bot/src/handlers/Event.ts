import { Client } from "discord.js"
import { readdirSync } from "fs"
import { join } from "path"
import { BotEvent } from "../types"
import log from "../lib/log";

module.exports = (client: Client) => {
    let eventsDir = join(__dirname, "../events")
    let eventCount = 0
    log('', false)
    log('&c:blue;&s:underscore;Events:', false)
    readdirSync(eventsDir).forEach(file => {
        eventCount++
        if (!file.endsWith('.ts')) return log(`&c:red;Found impostor file '${file}'`)
        let event: BotEvent = require(`${eventsDir}/${file}`).default
        event.once ?
            client.once(event.name, (...args) => event.execute(...args))
            :
            client.on(event.name, (...args) => event.execute(...args))

        log(`&c:yellow;Event '${event.name}' registered`)
    })
    log('', false)
    log(`&c:green;Successfully loaded &bg:green;&c:black; ${eventCount} &s:reset;&c:green; event${eventCount == 1? '': 's'}`)
}

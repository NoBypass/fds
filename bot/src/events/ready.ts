import { Client } from "discord.js"
import { BotEvent } from "../types"
import log from "../lib/log";

const event : BotEvent = {
    name: "ready",
    once: true,
    execute: () => {
        log('&c:black;&bg:blue; Bot started up ', false)
    }
}

export default event
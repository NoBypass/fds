import { BotEvent } from "../types"
import { Message } from "discord.js"

const event : BotEvent = {
    name: "messageCreate",
    once: true,
    execute: (message: Message) => {

    }
}

export default event
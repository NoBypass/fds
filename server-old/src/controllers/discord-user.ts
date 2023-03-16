import { DiscordUser, discordUser } from "../schemas/database/discord-user"

export default async function DiscordUser(param: string, method: string, body?: any) {
    switch (method) {
        case 'POST': {
            try {
                await discordUser.create<DiscordUser>({ id: param, ...body})
                return { success: true }
            } catch (e) {
                return {
                    success: false,
                    error: e
                }
            }
        }
        case 'GET': {
            try {
                return {
                    success: true,
                    ...await discordUser.findOne({ id: param })
                }
            } catch (e) {
                return {
                    success: false,
                    error: e
                }
            }
        }
        case 'PUT': {
            try {
                const { xp, claimedDaily, vc, msg } = body
                const previous = await discordUser.findOne({ id: param })
                if (previous == null) return { success: false, error: 'Could not find element to update' }
                const res = discordUser.findOneAndUpdate({ id: param }, {
                    id: param,
                    uuid: previous.uuid,
                    $inc: {
                        xp: xp? xp: 0,
                        dailiesClaimed: claimedDaily? 1: 0,
                        minutesSpentInVc: vc? 1: 0,
                        messagesSent: msg? 1: 0,
                    },
                    lastDailyClaimed: claimedDaily? new Date().getTime(): previous.lastDailyClaimed,
                    dailiesStreak: claimedDaily? previous.dailiesStreak + 1: previous.lastDailyClaimed + 1000 * 60 * 60 * 24 < new Date().getTime()? 0: previous.lastDailyClaimed
                })

                return {
                    success: true,
                    ...res
                }
            } catch (e) {
                return {
                    success: false,
                    error: e
                }
            }
        }
    }
}
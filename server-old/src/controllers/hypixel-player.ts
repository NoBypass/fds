import { hypixelPlayer } from "../schemas/database/hypixel-player"

export default async function HypixelPlayer(param: string, method: string): Promise<any> {
    const playerOld = await hypixelPlayer.findOne({ uuid: param })
    if (playerOld) {
        if (playerOld.player[0].timestamp + 1000 * 60 > new Date().getTime()) return {
            success: true,
            ...playerOld
        }

        return hypixelPlayer.findOneAndUpdate(
            { uuid: param },
            { $push: { player: await getHypixelData(param) } }
        )
    }
    const hypixelData = await getHypixelData(param)
    return hypixelPlayer.create({
        name: hypixelData.player.displayName,
        uuid: param,
        player: [
            { ...hypixelData }
        ]
    })
}

const getHypixelData = async (uuid: string) => {
    return await fetch(`https://api.hypixel.net/player?uuid=${uuid}?key=${process.env.HYPIXEL_API_KEY}`).then(r => r.json())
}
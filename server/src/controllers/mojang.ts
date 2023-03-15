import { mojang } from "../schemas/database/mojang";
import log from "../functions/log";

export default async function Mojang(param: string, method: string): Promise<any> {
    let mojangData = await mojang.findOne({ name: param })
    if (mojangData?.uuid != null) return {
        success: true,
        skin: mojangData.skin,
        name: mojangData.name,
        uuid: mojangData.uuid
    }

    mojangData = await fetch(`https://api.mojang.com/users/profiles/minecraft/${param}`)
        .then(r => r.json())
        .catch(e => {
            return {
                success: false,
                error: 'HypixelPlayer not found'
            }
        })
    if (mojangData?.id == null) return {
        success: false,
        error: 'HypixelPlayer not found'
    }

    const skin = await fetch(`https://crafatar.com/avatars/${mojangData.id}`)
        .then(async (response: Response) => {
            const arrayBuffer = await response.arrayBuffer()
            const imageBuffer = Buffer.from(arrayBuffer)
            return imageBuffer.toString('base64')
        })
        .catch(e => {
            log(`  -> Failed to get skin for ${(mojangData as any).id}`, 'c_red')
        })

    await mojang.create({
        uuid: mojangData.id,
        name: mojangData.name,
        skin: skin
    })
    return {
        success: true,
        uuid: mojangData.id,
        name: mojangData.name,
        skin: skin
    }
}
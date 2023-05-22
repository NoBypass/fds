const URI = process.env.API_URI || ''
const defaults = {
    method: 'post',
    headers: {
        'content-type': 'application/json'
    },
}

export const createUser = async (user: { id: number, tag: string, ign: string }) => {
      const hypixelPlayer = await fetch(URI, {
        ...defaults,
        body: `{
            "query": "{ hypixelPlayer(name: \"${user.ign}\") { name, stat_snapshots, guild_id } }"
        }`
    }).then(r => r.json())

}
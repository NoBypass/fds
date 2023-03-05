const url = 'http://localhost:5001'

//TODO give correct datatypes by making shared types between server and client

export const getMojangPlayer = async (ign: string): Promise<any> => {
    return get(`mojang/${ign}`)
}
export const getUser = async (uuid: string): Promise<any> => {
    return get(`user/${uuid}`)
}
export const createUser = async (body: any): Promise<any> => {
    return create(`user/${body.uuid}`, body)
}

const get = async (route: string): Promise<any> => {
    try { return await fetch(`${url}/${route}`).then(r => r.json())
    } catch (e) { return {
        success: false,
        error: 'GET request failed'
    }}
}
const create = async (route: string, body: unknown) => {
    try { return await fetch(`${url}/${route}`, {
        method: 'POST',
        body: JSON.stringify(body)
    }).then(r => r.json())
    } catch (e) { return {
        success: false,
        error: 'POST request failed'
    }}
}
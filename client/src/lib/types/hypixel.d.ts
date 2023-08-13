export {}

export interface Player {
    player: {
        id: string
        name: string
        uuid: string
    }
}

export type GraphQLRes <T> = {
    headers: {
        status: number
    }
    data: T
} | string
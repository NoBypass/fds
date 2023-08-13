export {}

export interface Player {
    id: string
    name: string
    uuid: string
}

export type GraphQLRes <T> = {
    headers: {
        status: number
        queryName: string
    }
    data: T
} | string
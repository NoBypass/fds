export type AccountWithName = {
    readonly name: string
}

export type ResError = {
    status: number
    msg: string
    error: string
}

export type Leaderboard = {
    rank: string
    country: string
    name: string
    stars: string
    fkdr: string
}

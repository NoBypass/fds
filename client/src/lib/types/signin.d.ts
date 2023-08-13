export type SigninInfo = {
    username: string
    password: string
    remember: boolean
}

export type SigninRes = {
    id: number
    username: string
    password: string
    token: string
    joined_at: string
}
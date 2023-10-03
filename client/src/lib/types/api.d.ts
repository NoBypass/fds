export type WsResponse<T> = {
    operationName: string
    data: T
    errors?: string[]
}

export type AccountWithName = {
    readonly name: string
}

export type ResError = {
    status: number
    msg: string
    error: string
}

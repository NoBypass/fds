export type WsResponse<T> = {
    operationName: string
    data: T
    errors?: string[]
}
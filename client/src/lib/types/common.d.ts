export type Point = {
    x: number,
    y: number
}

export type Listener = {
    run: (...args: any[]) => any
    event: string
}

export type FormField<T> = {
    val: T
    error: string
}

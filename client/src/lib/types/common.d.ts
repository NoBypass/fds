export type Point = {
    x: number,
    y: number
}

export type Listener = {
    run: (...args: any[]) => any
    event: string
}
import { writable } from 'svelte/store'

const curr = { x: 0, y: 0 }
export const mouse = writable(curr)

export const mouseStore = {
    update: (x: number, y: number) => {
        mouse.update(() => ({ x, y }))
        curr.x = x
        curr.y = y
    },
    subscribe: (run: (value: { x: number, y: number }) => void) => {
        mouse.subscribe(run)
    },
    intersects: (rect: DOMRect): boolean => {
        const { x, y } = curr
        return x >= rect.left
            && x <= rect.right
            && y >= rect.top
            && y <= rect.bottom
    }
}

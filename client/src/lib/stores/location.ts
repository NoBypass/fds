import { writable } from 'svelte/store'
import type { Listener } from '$lib/types/common'

const curr = { x: 0, y: 0 }
export const mouse = writable(curr)

const listeners: Listener[] = []

export const mouseStore = {
    move: (x: number, y: number) => {
        mouse.update(() => ({ x, y }))
        curr.x = x
        curr.y = y
    },
    click: () => {
        listeners.forEach(({ run, event }) => {
            if (event === 'click') run()
        })
    },
    onMove: (run: (value: { x: number, y: number }) => void) => {
        mouse.subscribe(run)
    },
    intersects: (rect: DOMRect): boolean => {
        const { x, y } = curr
        return x >= rect.left
            && x <= rect.right
            && y >= rect.top
            && y <= rect.bottom
    },
    onClick: (exec: () => void, rect?: DOMRect) => {
        const run = () => {
            if (rect && !mouseStore.intersects(rect)) return
            exec()
        }
        listeners.push({ run, event: 'click' })
    }
}

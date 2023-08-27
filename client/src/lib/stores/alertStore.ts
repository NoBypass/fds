import { writable } from 'svelte/store'
import type { Alert, ColorProp } from '$lib/types/components'
import { tweened } from 'svelte/motion'
import { cubicOut } from 'svelte/easing'

export const alerts = writable<Alert[]>([])
let sessionAlerts = 0

export const alertStore = {
    push: (message: string, color: ColorProp = 'info', icon = true) => {
        const id = sessionAlerts++
        const stage = tweened(0, { duration: 150, easing: cubicOut })
        stage.set(1).then(() => setTimeout(() => stage.set(0), 4750))

        stage.subscribe((value: number) => {
            alerts.update((alerts: Alert[]) =>
                alerts.map((alert: Alert) =>
                    alert.id === id ? { ...alert, stage: value } : alert)
            )
        })

        alerts.update((alerts: Alert[]) => [...alerts, {
            id, message, color, icon, stage: 0
        }])

        setTimeout(() => {
            alerts.update((alerts: Alert[]) => alerts.slice(1))
        }, 5000)
    },
    subscribe: alerts.subscribe
}

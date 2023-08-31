import { writable } from 'svelte/store'
import type { Alert, ColorProp } from '$lib/types/components'
import { tweened } from 'svelte/motion'
import { cubicOut } from 'svelte/easing'

export const alerts = writable<Alert[]>([])


export const alertStore = {
    push: (message: string, color: ColorProp = 'info', icon = true) => {
        const id = new Date().getTime()
        const tweenedStage = tweened(0, { duration: 150, easing: cubicOut })
        tweenedStage.set(1).then(() => setTimeout(() => tweenedStage.set(0), 4750))

        const timeout = setTimeout(() => {
            alerts.update((alerts: Alert[]) => alerts.slice(1))
        }, 5000)

        tweenedStage.subscribe((value: number) => {
            alerts.update((alerts: Alert[]) =>
                alerts.map((alert: Alert) =>
                    alert.id === id ? { ...alert, stage: value } : alert)
            )
        })

        alerts.update((alerts: Alert[]) => [...alerts, {
            id, message, color, icon, stage: 0, timeout, tweenedStage
        }])
    },
    remove: (alert: Alert) => {
        clearTimeout(alert.timeout)
        alert.tweenedStage.set(0)
        setTimeout(() => {
            alerts.update((alerts: Alert[]) => alerts.filter((a: Alert) => a.id !== alert.id))
        }, 250)
    }
}

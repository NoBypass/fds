import { readable } from 'svelte/store'
import WebSocket_2 from 'vite'
import WebSocket = WebSocket_2.WebSocket;
import type { WsResponse } from '$lib/types/api'

const instance = new WebSocket('ws://localhost:8080/ws') // TODO: use env variable

const api = {
    query: async <T> (query: string, operationName?: string) => {
        try {
            return await newWsRequest<T>(instance, query, operationName || '')
        } catch (error) {
            console.error('WebSocket request error:', error)
            throw error
        }
    },
    instance
}

const newWsRequest = <T>(
    ws: WebSocket,
    query: string,
    opName: string,
    timeoutDuration = 5000,
): Promise<T> => {
    return new Promise((resolve, reject) => {
        const timeout = setTimeout(() => {
            reject(new Error('WebSocket request timed out'))
        }, timeoutDuration)

        const handleMessage = (event: WebSocket.MessageEvent) => {
            try {
                const response: WsResponse<T> = JSON.parse(event.data as string)
                if (response.operationName === opName) {
                    clearTimeout(timeout)
                    if (response.errors) reject(new Error('WebSocket request failed'))
                    else resolve(response.data)
                    ws.removeEventListener('message', handleMessage)
                }
            } catch (error) {
                reject(error)
            }
        }

        ws.addEventListener('message', handleMessage)

        ws.send(JSON.stringify({
            operationName: opName || '',
            query: query.replace(/\s+/g, ''),
            variables: {},
        }))
    })
}


export const ws = readable(api)
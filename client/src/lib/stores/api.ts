import { readable } from 'svelte/store'
import type { WsResponse } from '$lib/types/api'

type Data = string | Buffer | ArrayBuffer | Buffer[]
interface MessageEvent {
    data: Data
    type: string
    target: WebSocket
}

const schema = (instance: WebSocket) => {
    return {
        query: async <T> (query: string, operationName?: string) => {
            try {
                return await newWsRequest<T>(instance, query, operationName || '')
            } catch (error) {
                console.error('WebSocket request error:', error)
                throw error
            }
        },
    }
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

        const handleMessage = (event: MessageEvent) => {
            try {
                const response: WsResponse<T> = JSON.parse(event.data as string)
                if (response.operationName === opName) {
                    clearTimeout(timeout)
                    if (response.errors) reject(new Error('WebSocket request failed'))
                    else resolve(response.data)
                    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
                    // @ts-ignore
                    ws.removeEventListener('message', handleMessage)
                }
            } catch (error) {
                reject(error)
            }
        }

        // eslint-disable-next-line @typescript-eslint/ban-ts-comment
        // @ts-ignore
        ws.addEventListener('message', handleMessage)

        ws.send(JSON.stringify({
            operationName: opName || '',
            query: query.replace(/\s+/g, ''),
            variables: {},
        }))
    })
}

export const api = readable({
    connect: (instance: WebSocket) => {
        return schema(instance)
    },
    graphql: {
        query: async <T>(query: string, operationName?: string) => {
            const res = await fetch('http://localhost:8080/graphql', { // TODO: use env variable
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    operationName: operationName || '',
                    query: query.replace(/\s+/g, ''),
                    variables: {},
                }),
            })
            return (await res.json() as {
                data: {
                    [key: string]: T
                }}).data[query.split('{')[1].split('(')[0].trim()]
        },
    },
})
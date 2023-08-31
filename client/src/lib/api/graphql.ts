import type { GraphQLRes } from '$lib/types/hypixel'
import { alertStore } from '$lib/stores/alertStore'

export const makeGraphQLRequest = async <T> (query: string, variables?: any): Promise<GraphQLRes<Partial<T>>> => {
    try {
        query = query.trim().replace(/\s+/g, ' ').replace('\n', '')
        const res = await fetch('http://localhost:8080/graphql', { // TODO: use env variable
            method: 'POST',
            body: JSON.stringify({
                query,
                operationName: '',
                variables: variables || {}
            })
        })

        const queryName = query.split('{')[1].split('(')[0].trim()
        return {
            headers: {
                status: res.status,
                queryName
            },
            data: (await res.json()).data[queryName] as Partial<T>
        }
    } catch (error) {
        alertStore.push('Request to server failed, please try again later', 'danger')
        return 'GraphQL request failed: ' + error
    }
}
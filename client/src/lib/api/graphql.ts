import type { GraphQLRes } from '$lib/types/hypixel'

export const makeGraphQLRequest = async <T> (query: string, variables?: any): Promise<GraphQLRes<Partial<T>>> => {
    try {
        query = query.trim().replace(/\s+/g, ' ').replace('\n', '')
        const res = await fetch('http://localhost:8080/graphql', { // TODO: dotenv
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
        return 'GraphQL request failed: ' + error
    }
}
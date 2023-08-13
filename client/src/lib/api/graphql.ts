import type { GraphQLRes } from '$lib/types/hypixel'

export const makeGraphQLRequest = async <T> (query: string, variables?: any): Promise<GraphQLRes<Partial<T>>> => {
    try {
        const res = await fetch('http://localhost:8080/graphql', { // TODO: dotenv
            method: 'POST',
            body: JSON.stringify({
                query,
                operationName: '',
                variables: variables || {}
            })
        })

        return {
            headers: {
                status: res.status
            },
            data: await res.json() as Partial<T>
        }
    } catch (error) {
        return 'GraphQL request failed: ' + error
    }
}
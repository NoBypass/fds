import type { AccountWithName } from '$lib/types/api'
import query from '$lib/api/common'

const api = {
    account: {
        name: async (name: string) => {
            return query<AccountWithName>({
                query: `
                    query ($name: String!) {
                        account(name: $name) {
                            name
                        }
                    }
                `,
                variables: {
                    name,
                }
            })
        }
    },
    auth: async ({ name, password, remember }: { name: string, password: string, remember: boolean }) => {
        return query<{
            readonly token: string,
            readonly account: AccountWithName
        }>({
            query: `
                query ($name: String!, $password: String!, $remember: Boolean!) {
                    signin(name: $name, password: $password, remember: $remember) {
                        token, account { name }
                    }
                }
            `,
            variables: {
                name,
                password,
                remember,
            }
        })
    },
    player: {
        name: async (name: string) => {
            return query<{
                readonly name: string
            }>({
                query: `
                    query ($name: String!) {
                        player(name: $name) {
                            name
                        }
                    }
                `,
                variables: {
                    name,
                }
            })
        }
    }
}

export default api

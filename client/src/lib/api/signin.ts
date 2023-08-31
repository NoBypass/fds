import type { SigninInfo, SigninRes } from '$lib/types/signin'
import { makeGraphQLRequest } from '$lib/api/graphql'

export const signin = async (info: SigninInfo) => {
    return makeGraphQLRequest<SigninRes>(`mutation {
        signin(username: "${info.username}", password: "${info.password}", remember: ${info.remember}) {
            token username
        }
    }`)
}

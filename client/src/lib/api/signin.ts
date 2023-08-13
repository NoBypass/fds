import type { SigninInfo, SigninRes } from '$lib/types/signin'
import { makeGraphQLRequest } from '$lib/api/graphql'

export const signin = async (info: SigninInfo) => {
    return makeGraphQLRequest<SigninRes>(`mutation {
        signin(username: $username, password: $password, remember: $remember) {
            token
        }
    }`, info)
}
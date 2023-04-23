import {gql} from 'apollo-server-express'
import {getTypeDefStrings} from "./resolver"

const {mutations, queries, types} = getTypeDefStrings()
const completeQueries = `type Query {${queries.map((q: string) => q)}}`
const completeMutations = `type Mutation {${mutations.map((m: string) => m)}}`

export const typeDefs = gql`
    ${types.map((t: string) => t)}

    ${completeQueries}

    ${completeMutations}
`
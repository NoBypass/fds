import {Redis} from "ioredis"
import Resolver from "./_Resolver"
import {GraphQLProperty} from "../../types/GraphQLProperty"

export default class Getter extends Resolver {
    resolverType: 'query' = 'query'

    constructor(tableName: string, inputSchema: GraphQLProperty, middleware?: (dto: any) => any) {
        super(tableName, inputSchema, middleware)
    }

    resolver = {
        [this.tableName]: async (
            _: void,
            {id}: {id: string},
            {redis}: { redis: Redis }
        ): Promise<string | null> => {
            if (id == null) throw new Error('Field "id" not found')
            if (this.middleware != null) id = await this.middleware(id)
            const discordUser = await redis.get(`discordUser:${id}`)
            return discordUser ? JSON.parse(discordUser) : null
        }
    }

    getResolverAsGraphQLString(): string {
        return `${this.tableName}(id: ID!): ${this.tableName.replace(/^\w/, (c) => c.toUpperCase())}`
    }
}

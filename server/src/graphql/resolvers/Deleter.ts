import {Redis} from "ioredis"
import Resolver from "./_Resolver"
import {GraphQLProperty} from "../../types/GraphQLProperty";
import {firstLetterUpperCase, getPropertiesAmount} from "../../lib/common";

export default class Deleter extends Resolver {
    resolverType: 'mutation' = 'mutation'

    constructor(tableName: string, inputSchema: GraphQLProperty, middleware?: (dto: any) => any) {
        super(tableName, inputSchema, middleware)
    }

    resolver = {
        [`delete${this.tableName.replace(/^\w/, (c) => c.toUpperCase())}`]: async (
            _: void,
            inputData: any,
            {redis}: { redis: Redis }
        ): Promise<string | null> => {
            if (getPropertiesAmount(inputData) != 1) return null
            let id = inputData[Object.keys(this.inputSchema)[0]]
            if (this.middleware != null) id = await this.middleware(id)
            const response = await redis.get(`${this.tableName}:${id}`)
            if (!response) return null
            await redis.multi()
                .lrem(`${this.tableName}s`, 1, response)
                .del(`${this.tableName}:${id}`)
                .exec()
            return id
        }
    }

    getResolverAsGraphQLString(): string {
        return `delete${firstLetterUpperCase(this.tableName)}(id: ID!): ID!`
    }
}
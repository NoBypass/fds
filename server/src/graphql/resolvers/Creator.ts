import {Redis} from "ioredis"
import Resolver from "./_Resolver"
import {GraphQLProperty} from "../../types/GraphQLProperty";
import {firstLetterUpperCase, generateUUID} from "../../lib/common";

export default class Creator extends Resolver {
    resolverType: 'mutation' = 'mutation'

    constructor(tableName: string, inputSchema: GraphQLProperty, middleware?: (dto: any) => any) {
        super(tableName, inputSchema, middleware)
    }

    resolver = {
        [`create${this.tableName.replace(/^\w/, (c) => c.toUpperCase())}`]: async (
            _: void,
            inputData: any,
            {redis}: { redis: Redis }
        ): Promise<string | null> => {
            const id = generateUUID()
            if (this.middleware != null) inputData = await this.middleware(inputData)
            await redis.multi()
                .lpush(`${this.tableName}s`, inputData)
                .set(`${this.tableName}:${id}`, inputData)
                .exec()
            return inputData
        }
    }

    getResolverAsGraphQLString(): string {
        let params = ''
        for (let [key, value] of Object.entries(this.inputSchema)) {
            params += `${key}: ${value.replace('!', '')}, `
        }
        return `create${firstLetterUpperCase(this.tableName)}(${params}): ${firstLetterUpperCase(this.tableName)}!`
    }
}
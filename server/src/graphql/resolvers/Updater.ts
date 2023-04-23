import {Redis} from "ioredis"
import Resolver from "./_Resolver"
import {GraphQLProperty} from "../../types/GraphQLProperty";
import {firstLetterUpperCase} from "../../lib/common";

export default class Updater extends Resolver {
    resolverType: 'mutation' = 'mutation'

    constructor(tableName: string, inputSchema: GraphQLProperty, middleware?: (dto: any) => any) {
        super(tableName, inputSchema, middleware)
    }

    resolver = {
        [`update${this.tableName.replace(/^\w/, (c) => c.toUpperCase())}`]: async (
            _: void,
            inputData: any,
            {redis}: { redis: Redis }
        ): Promise<string | null> => {
            if (inputData.id == null) throw new Error('Input data did not contain field "id"')
            const id = inputData.id
            if (this.middleware != null) inputData = await this.middleware(inputData)
            const response = await redis.get(`${this.tableName}:${id}`)
            if (!response) return null
            const updatedData = {
                ...JSON.parse(response),
                ...inputData
            }
            await redis.multi()
                .lrem(`${this.tableName}s`, 1, response)
                .lpush(`${this.tableName}s`, updatedData)
                .set(`${this.tableName}:${id}`, updatedData)
                .exec()
            return updatedData
        }
    }

    getResolverAsGraphQLString(): string {
        let params = ''
        for (let [key, value] of Object.entries(this.inputSchema)) {
            params += `${key}: ${value.replace('!', '')}, `
        }
        return `update${firstLetterUpperCase(this.tableName)}(${params}): ${firstLetterUpperCase(this.tableName)}!`
    }
}
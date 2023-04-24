import {Redis} from "ioredis"
import {GraphQLProperty} from "../../types/GraphQLProperty"

export default abstract class Resolver {
    readonly tableName: string
    middleware: null | ((dto: any) => Promise<any>) = null
    inputSchema: GraphQLProperty

    abstract resolverType: 'mutation' | 'query'
    abstract resolver:  { [p: string]: (_: void, input: any, {redis}: {redis: Redis}) => Promise<string | null> }

    protected constructor(tableName: string, inputSchema: GraphQLProperty, middleware?: (dto: any) => any) {
        this.tableName = tableName
        this.inputSchema = inputSchema
        this.middleware = middleware == null ? null : middleware
    }

    get(): any {
        return this.resolver
    }

    getResolverType(): 'mutation' | 'query' {
        return this.resolverType
    }

    abstract getResolverAsGraphQLString(): string
}
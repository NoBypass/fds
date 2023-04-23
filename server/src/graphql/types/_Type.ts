import {GraphQLProperty} from "../../types/GraphQLProperty";

export default abstract class GraphQL {
    abstract tableName: string
    abstract properties: GraphQLProperty
    abstract resolvers: any[]

    getPropertiesAsGraphQLStrings(): string[] {
        let arr: string[] = []
        for (const [key, val] of Object.entries(this.properties)) arr.push(`${key}: ${val}`)
        return arr
    }

    getTableName(): string {
        return this.tableName
    }

    getResolvers(): any[] {
        return this.resolvers
    }
}
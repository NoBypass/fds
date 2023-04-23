import GraphQL from "./_Type"
import {GraphQLProperty, GraphQLTypes} from "../../types/GraphQLProperty"
import Deleter from "../resolvers/Deleter"
import Updater from "../resolvers/Updater"
import Creator from "../resolvers/Creator"
import Getter from "../resolvers/Getter"

const {
    ID,
    Int,
    ReqID,
    ReqInt,
    String,
    ReqString,
    ReqBoolean,
    Boolean} = GraphQLTypes

export class HypixelPlayer extends GraphQL {
    tableName = 'hypixelPlayer'

    properties: GraphQLProperty = {
        id: ReqID,
        name: ReqString,
        description: String,
        motd: String,
        level: ReqInt,
        xp: ReqInt,
        created_at: ReqInt,
        is_tracked: ReqBoolean
    }

    resolvers = [
        new Deleter(this.tableName, {id: ReqID}),
        new Updater(this.tableName, {
            id: ReqID,
            name: String,
            description: String,
            motd: String,
            is_tracked: Boolean
        }),
        new Creator(this.tableName, {
            name: ReqString,
            description: String,
            motd: String,
            is_tracked: ReqBoolean
        }),
        new Getter(this.tableName, {id: ReqID})
    ]
}
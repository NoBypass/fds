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
    ReqBoolean,
    Boolean} = GraphQLTypes

export class HypixelPlayer extends GraphQL {
    tableName = 'hypixelPlayer'

    properties: GraphQLProperty = {
        id: ReqID,
        mojang_user_id: ReqID,
        guild_id: ReqID,
        latest_player_stats: String,
        latest_lookup_at: Int,
        registration_at: ReqInt,
        is_tracked: ReqBoolean
    }

    resolvers = [
        new Deleter(this.tableName, {id: ReqID}),
        new Updater(this.tableName, {
            id: ReqID,
            guild_id: ID,
            latest_player_stats: String,
            latest_lookup_at: Int,
            is_tracked: Boolean
        }),
        new Creator(this.tableName, {
            mojang_user_id: ReqID,
            latest_player_stats: String,
            is_tracked: Boolean
        }),
        new Getter(this.tableName, {id: ReqID})
    ]
}
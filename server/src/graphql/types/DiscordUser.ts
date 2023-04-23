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
    ReqInt} = GraphQLTypes

export class DiscordUser extends GraphQL {
    tableName = 'discordUser'

    properties: GraphQLProperty = {
        id: ReqID,
        hypixel_player_id: ReqID,
        uuid: ReqInt,
        level: ReqInt,
        overflow_xp: ReqInt,
        dailies_streak: ReqInt,
        xp_from_dailies: ReqInt,
        last_daily_claimed: ReqInt,
        minutes_spent_in_vc: ReqInt,
        messages_sent: ReqInt
    }

    resolvers = [
        new Deleter(this.tableName, {id: ReqID}),
        new Updater(this.tableName, {
            id: ReqID,
            hypixel_player_id: ID,
            level: Int,
            overflow_xp: Int,
            dailies_streak: Int,
            xp_from_dailies: Int,
            last_daily_claimed: Int,
            minutes_spent_in_vc: Int,
            messages_sent: Int
        }),
        new Creator(this.tableName, {
            hypixel_player_id: ReqID,
            uuid: ReqInt
        }),
        new Getter(this.tableName, {id: ReqID})
    ]
}
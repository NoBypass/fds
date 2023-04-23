import GraphQL from "./_Type"
import {GraphQLProperty, GraphQLTypes} from "../../types/GraphQLProperty"
import Deleter from "../resolvers/Deleter"
import Updater from "../resolvers/Updater"
import Creator from "../resolvers/Creator"
import Getter from "../resolvers/Getter"

const {
    ID,
    ReqID,
    String} = GraphQLTypes

export class MinecraftSkin extends GraphQL {
    tableName = 'minecraftSkin'

    properties: GraphQLProperty = {
        id: ReqID,
        skin_base64: String
    }

    resolvers = [
        new Deleter(this.tableName, {id: ReqID}),
        new Creator(this.tableName, {
            mojang_user_id: ID
        }),
        new Getter(this.tableName, {id: ReqID})
    ]
}
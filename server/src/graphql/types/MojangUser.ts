import GraphQL from "./_Type"
import {GraphQLProperty, GraphQLTypes} from "../../types/GraphQLProperty"
import Deleter from "../resolvers/Deleter"
import Updater from "../resolvers/Updater"
import Creator from "../resolvers/Creator"
import Getter from "../resolvers/Getter"

const {
    ID,
    ReqID,
    String,
    ReqString} = GraphQLTypes

export class HypixelPlayer extends GraphQL {
    tableName = 'mojangUser'

    properties: GraphQLProperty = {
        id: ReqID,
        minecraft_skin_id: ReqID,
        name: ReqString,
        uuid: ReqString
    }

    resolvers = [
        new Deleter(this.tableName, {id: ReqID}),
        new Updater(this.tableName, {
            id: ReqID,
            name: String,
            minecraft_skin_id: ID
        }),
        new Creator(this.tableName, {
            name: ReqString
        }),
        new Getter(this.tableName, {id: ReqID})
    ]
}
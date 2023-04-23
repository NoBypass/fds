export enum GraphQLTypes {
    Int = 'Int',
    ReqInt = 'Int!',
    String = 'String',
    ReqString = 'String!',
    ID = 'ID',
    ReqID = 'ID!',
    Float = 'Float',
    ReqFloat = 'Float!',
    Boolean = 'Boolean',
    ReqBoolean = 'Boolean!'
}

export type GraphQLProperty = {
    [key: string]: GraphQLTypes
}
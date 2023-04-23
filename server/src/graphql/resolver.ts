import * as fs from 'fs'
import * as path from 'path'
import {encode, firstLetterUpperCase, generateUUID, shuffle} from "../lib/common";

export const resolvers = {
    Query: {},
    Mutation: {}
}

const typeDefStrings: {
    types: string[]
    queries: string[]
    mutations: string[]
} = {
    types: [],
    queries: [],
    mutations: []
}

const cycleThroughFiles = () => {
    const folderPath = path.resolve(__dirname, './types')
    try {
        const files = fs.readdirSync(folderPath);
        files.forEach((file) => {
            const filePath = path.join(folderPath, file)
            const splitPath = filePath.split('/')
            if (!filePath.endsWith('.ts') || splitPath[splitPath.length - 1].startsWith('_')) return
            const module = require(filePath)
            const instance = new module[path.basename(filePath, '.ts')]()
            const localResolvers = instance.getResolvers()
            localResolvers.forEach((r: any) => {
                const resolverType = r.getResolverType()
                if (resolverType == 'mutation') {
                    resolvers.Mutation = {...resolvers.Mutation, ...r.get()}
                    typeDefStrings.mutations.push(r.getResolverAsGraphQLString())
                } else if (resolverType == 'query') {
                    resolvers.Query = {...resolvers.Query, ...r.get()}
                    typeDefStrings.queries.push(r.getResolverAsGraphQLString())
                }
            })
            typeDefStrings.types.push(`type ${firstLetterUpperCase(instance.getTableName())} {${instance.getPropertiesAsGraphQLStrings().map((p: string) => `${p}, `)}}`)
        })
    } catch (err) {
        console.error(err)
    }
}

export const getTypeDefStrings = () => {
    cycleThroughFiles()
    console.log(typeDefStrings, resolvers)
    return typeDefStrings
}
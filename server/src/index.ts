import express from 'express'
import { ApolloServer } from 'apollo-server-express'
import { typeDefs } from './graphql/schema'
import { resolvers } from './graphql/resolver'
import { redis } from './config/redis'

const app = express()
const port = 6419

const server = new ApolloServer({
    typeDefs,
    resolvers,
    context: () => ({ redis }),
})

async function startServer() {
    await server.start()
    server.applyMiddleware({ app })
}
startServer().then()

app.listen({ port }, () => {
    console.log(`Server ready at http://localhost:${port}${server.graphqlPath}`);
})
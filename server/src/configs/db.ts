require('dotenv').config({path:__dirname+'/../.env'})

export const dbConfig = {
    mongo: {
        url: process.env.MONGO_URI || ''
    },
    server: {
        port: process.env.MONGO_PORT ? Number(process.env.MONGO_PORT) : 5002
    }
}
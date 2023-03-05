import express from 'express'
import connectToDb from "./src/functions/database-connect"
import { dbConfig } from "./src/configs/db"
import log from "./src/functions/log"
import getRoutes from "./src/functions/get-routes"
import { routes } from "./src/schemas/routes"
import listenOnRoute from "./src/functions/listen-on-route"
const cors = require('cors')
require('dotenv').config({path:__dirname+'/./src/.env'})

const app = express()

const onStart = async () => {
    await connectToDb(dbConfig)
}
onStart()
app.use(cors())
app.use(express.json())

log('FDS Hub server', undefined, ['bg_blue', 'c_black'], false)
console.log('')
log('Routes:', 'c_blue')

const routeArray = getRoutes(routes)
routeArray.forEach(route => {
    log(`Listening on ${route.route} (${route.method})`, 'c_green')
    listenOnRoute(route, app)
})

app.listen(process.env.PORT, () => {
    console.log('')
    log('Listening on port ' + process.env.PORT, undefined, ['bg_green', 'c_black'])
})


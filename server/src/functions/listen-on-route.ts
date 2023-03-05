import { FinalRoute } from "../schemas/final-route"
import { Express } from "express"
import log from "./log"

export default function listenOnRoute(routeObj: FinalRoute, app: Express): void {
    const { route, method } = routeObj
    if (method === 'all') {
        ['put', 'post', 'get'].forEach((m) => {
            (app.route(route) as any)[m](async (req: any, res: any) => {
                await handleRequest(routeObj, req, res, m.toUpperCase())
            })
        })
    } else {
        app.route(route)[method](async (req, res) => {
            await handleRequest(routeObj, req, res, method.toUpperCase())
        })
    }
}

async function handleRequest(routeObj: FinalRoute, req: any, res: any, method: string) {
    res.setHeader('Access-Control-Allow-Origin', '*')
    let processStart = new Date().getTime()
    log(`'${method}' request on ${routeObj.route.split(':')[0] + req.params.param}`, 'c_yellow')

    let result
    if (req.body == null) result = await routeObj.controller(req.params.param, method)
    else result = await routeObj.controller(req.params.param, method, req.body)
    res.json({
        ...result
    })
    log(`   -> Process ended after ${new Date().getTime() - processStart}ms`, `c_${result.success ? 'green' : 'red'}`)
}
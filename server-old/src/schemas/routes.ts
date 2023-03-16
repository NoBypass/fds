import User from "../controllers/user"
import Session from "../controllers/session"
import { Route } from "./route"
import HypixelPlayer from "../controllers/hypixel-player"
import Mojang from "../controllers/mojang"

export const routes: Route[] = [
    { name: 'user', method: 'all', controller: User },
    { name: 'discord', routes: [
            { name: 'user', method: 'all', controller: User },
        ] },
    { name: 'session', method: 'get', controller: Session },
    { name: 'hypixel', routes: [
            { name: 'player', method: 'get', controller: HypixelPlayer },
        ] },
    { name: 'mojang', method: 'get', controller: Mojang }
]
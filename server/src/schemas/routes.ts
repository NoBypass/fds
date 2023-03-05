import User from "../controllers/user"
import Session from "../controllers/session"
import { Route } from "./route"
import Player from "../controllers/player"
import Mojang from "../controllers/mojang"

export const routes: Route[] = [
    { name: 'user', method: 'all', controller: User },
    { name: 'session', method: 'get', controller: Session },
    { name: 'hypixel', routes: [
            { name: 'player', method: 'get', controller: Player },
        ] },
    { name: 'mojang', method: 'get', controller: Mojang }
]
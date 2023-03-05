import { User, user } from "../schemas/database/user"
import jwt from 'jsonwebtoken'


export default async function User(param: string, method: string, body?: any): Promise<any> { //TODO add correct body type

    switch (method) {
        case 'GET': {
            const userData = await user.findOne({ uuid: param }).select({ _id: 0 })
            if (userData == null) return {
                success: false,
                error: 'Couldn\'t find user'
            }
            let res: { [key: string]: any } = {
                success: true,
            }
            for (let [key, value] of Object.entries(userData)) {
                res[key] = value
            }
            return res
        }
        case 'POST': {
            try {
                const secret: string = process.env.JWT_SECRET || ''
                const payload = {
                    uuid: body.uuid,
                    expires_at: body.stayLogged? -1: new Date().getTime() / 1000 + 60 * 60 * 24 * 14,
                }
                await user.create<User>({
                    uuid: body.uuid,
                    password: body.password,
                    discord: body.discord,
                    registrationDate: new Date().getTime(),
                    confirmed: false, //TODO discord confirmation
                    settings: {} //TODO settings
                })
                const token = jwt.sign(payload, secret)
                return {
                    success: true,
                    token,
                }
            } catch (e) {
                return {
                    success: false,
                    error: 'Couldn\'t create user'
                }
            }
        }
    }
}
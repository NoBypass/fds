import mongoose from "mongoose"
import {DbConfig} from "../schemas/db-config";

export default async function connectToDb(config: DbConfig):Promise<void> {
    mongoose.set('strictQuery', false)
    await mongoose.connect(
        config.mongo.url || '',
        {
            keepAlive: true
        })
}
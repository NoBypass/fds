import mongoose, { Schema, InferSchemaType  } from 'mongoose'

const schema = new Schema({
    name: { type: String, required: true },
    uuid: { type: String, required: true },
    skin: { type: String, required: true }
})

export type Mojang = InferSchemaType<typeof schema>
export const mojang = mongoose.model<Mojang>('Mojang', schema)
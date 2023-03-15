import mongoose, { Schema, InferSchemaType  } from 'mongoose'

const schema = new Schema({
    name: { type: String, required: true },
    uuid: { type: String, required: true },
    player: [
        {
            timestamp: { type: Number, required: true },
            data: { type: Object, required: true }
        }
    ]
})

export type HypixelPlayer = InferSchemaType<typeof schema>
export const hypixelPlayer = mongoose.model<HypixelPlayer>('HypixelPlayer', schema)
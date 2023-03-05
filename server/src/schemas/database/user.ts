import mongoose, { Schema, InferSchemaType  } from 'mongoose'

const schema = new Schema({
    uuid: { required: true, type: String },
    password: { required: true, type: String },
    discord: { required: true, type: String },
    registrationDate: { required: true, type: Number },
    confirmed: { required: true, type: Boolean },
    settings: { required: false, type: Object }
})

export type User = InferSchemaType<typeof schema>
export const user = mongoose.model<User>('User', schema)
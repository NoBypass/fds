import mongoose, { Schema, InferSchemaType  } from 'mongoose'

const schema = new Schema({
    id: { required: true, type: String },
    uuid: { required: true, type: String },
    xp: { required: true, type: Number },
    dailiesClaimed: { required: true, type: Number },
    minutesSpentInVc: { required: true, type: Number },
    messagesSent: { required: true, type: Number },
    dailiesStreak: { required: true, type: Number },
    lastDailyClaimed: { required: true, type: Number }
})

export type DiscordUser = InferSchemaType<typeof schema>
export const discordUser = mongoose.model<DiscordUser>('DiscordUser', schema)
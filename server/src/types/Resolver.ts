import {Redis} from "ioredis"

export type Resolver = { [p: string]: (_: void, input: any, {redis}: {redis: Redis}) => Promise<string | null> }
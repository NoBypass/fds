export type Route = {
    name: string
    } & (
        {
            method: 'get' | 'post' | 'put' | 'all'
            controller: (param: string, method: string, body?: any) => Promise<any>
            routes?: never
        } | {
            method?: never
            controller?: never
            routes: Route[]
    }
)
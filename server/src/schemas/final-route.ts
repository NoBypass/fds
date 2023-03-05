export type FinalRoute = {
    route: string
    method: 'get' | 'post' | 'put' | 'all'
    controller: (param: string, method: string, body?: any) => Promise<any>
}
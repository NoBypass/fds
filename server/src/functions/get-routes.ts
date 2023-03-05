import { FinalRoute } from "../schemas/final-route"
export default function getRoutes(routeSchema: any, currentRoute = ""): FinalRoute[] {
    let routes: FinalRoute[] = []

    routeSchema.forEach((route: any) => {
        let newRoute = currentRoute + "/" + route.name

        if (route.routes) routes = routes.concat(getRoutes(route.routes, newRoute))
        else routes.push({
            route: newRoute + '/:param',
            method: route.method,
            controller: route.controller
        })
    });
    return routes
}
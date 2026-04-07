import { RouteProps } from "react-router-dom";
import Authentication from "../../pages/AuthenticationPage/Authentication";

export enum AppRoutes {
    AUTH = 'auth',
}

export const RoutePath: Record<AppRoutes, string> = {
    [AppRoutes.AUTH]: `/auth`,
}

export const routeConfig: Record<AppRoutes, RouteProps> = {
    [AppRoutes.AUTH]: {
        path: RoutePath.auth,
        element: <Authentication />
    }
}
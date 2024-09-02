export interface GyeonggiBusRouteResponse {
    TBBMSROUTEM: {
        head: {
            list_total_count: number;
            RESULT: {
                CODE: string;
                MESSAGE: string;
            };
        },
        row: GyeonggiBusRouteData[];
    }
}

export interface GyeonggiBusRouteData {
    ROUTE_NM: string;
    ROUTE_ID: string;
}

export interface GyeonggiBusRouteErrorResponse {
    RESULT: {
        CODE: string;
        MESSAGE: string;
    }
}

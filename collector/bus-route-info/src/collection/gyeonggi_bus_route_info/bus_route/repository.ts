import { OriginalGyeonggiBusRoute } from "@prisma/client";

import { prisma } from "../../../database";
import { GyeonggiBusRouteData } from "./interface";

export const upsert = async (gyeonggiBusRouteData: GyeonggiBusRouteData): Promise<void> => {
    await prisma.originalGyeonggiBusRoute.upsert({
        where: {
            route_id: gyeonggiBusRouteData.ROUTE_ID,
        },
        update: {
            route_name: gyeonggiBusRouteData.ROUTE_NM,
        },
        create: {
            route_id: gyeonggiBusRouteData.ROUTE_ID,
            route_name: gyeonggiBusRouteData.ROUTE_NM,
        }
    });
}

export const findMany = async (): Promise<OriginalGyeonggiBusRoute[]> => {
    return await prisma.originalGyeonggiBusRoute.findMany();
}

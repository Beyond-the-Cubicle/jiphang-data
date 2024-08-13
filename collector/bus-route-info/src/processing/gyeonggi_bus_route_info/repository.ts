import { prisma } from "../../database";
import { OriginalGyeonggiBusRouteInfo, BusRouteInfo } from "@prisma/client";

export const findManyOriginalGyeonggiBusRouteInfo = async (): Promise<OriginalGyeonggiBusRouteInfo[]> => {
    return await prisma.originalGyeonggiBusRouteInfo.findMany();
}

export const upsert = async (busRouteInfo: BusRouteInfo): Promise<void> => {
    await prisma.busRouteInfo.upsert({ 
        where: {
            route_id: busRouteInfo.route_id,
        },
        update: {
            route_name: busRouteInfo.route_name,
            route_type: busRouteInfo.route_type,
            start_station_name: busRouteInfo.start_station_name,
            end_station_name: busRouteInfo.end_station_name,
            maximum_dispatch_interval: busRouteInfo.maximum_dispatch_interval,
            minimum_dispatch_interval: busRouteInfo.minimum_dispatch_interval,
            start_station_first_time: busRouteInfo.start_station_first_time,
            start_station_last_time: busRouteInfo.start_station_last_time,
            start_station_low_first_time: busRouteInfo.start_station_low_first_time,
            start_station_low_last_time: busRouteInfo.start_station_low_last_time,
            end_station_first_time: busRouteInfo.end_station_first_time,
            end_station_last_time: busRouteInfo.end_station_last_time,
        },
        create: {
            route_id: busRouteInfo.route_id,
            route_name: busRouteInfo.route_name,
            route_type: busRouteInfo.route_type,
            start_station_name: busRouteInfo.start_station_name,
            end_station_name: busRouteInfo.end_station_name,
            maximum_dispatch_interval: busRouteInfo.maximum_dispatch_interval,
            minimum_dispatch_interval: busRouteInfo.minimum_dispatch_interval,
            start_station_first_time: busRouteInfo.start_station_first_time,
            start_station_last_time: busRouteInfo.start_station_last_time,
            start_station_low_first_time: busRouteInfo.start_station_low_first_time,
            start_station_low_last_time: busRouteInfo.start_station_low_last_time,
            end_station_first_time: busRouteInfo.end_station_first_time,
            end_station_last_time: busRouteInfo.end_station_last_time,
        }
    });
}

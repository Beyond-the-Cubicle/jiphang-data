import dayjs from "dayjs";
import { OriginalSeoulBusRouteInfo } from "@prisma/client";
import { findManyOriginalSeoulBusRouteInfo, upsert } from "./repository";

export const processSeoulBusRouteInfo = async (): Promise<void> => {
    const seoulBusRouteInfoList: OriginalSeoulBusRouteInfo[] = await findManyOriginalSeoulBusRouteInfo();
    seoulBusRouteInfoList
        .filter((seoulBusRouteInfo) => seoulBusRouteInfo.bus_route_nm
            && seoulBusRouteInfo.route_type
            && seoulBusRouteInfo.st_station_nm
            && seoulBusRouteInfo.ed_station_nm
            && seoulBusRouteInfo.term
        )
        .map((seoulBusRouteInfo) => {
            upsert({
                route_id: seoulBusRouteInfo.bus_route_id,
                route_region: 'SEOUL',
                route_name: seoulBusRouteInfo.bus_route_nm!,
                route_type: parseRouteType(seoulBusRouteInfo.route_type!),
                start_station_name: seoulBusRouteInfo.st_station_nm!,
                end_station_name: seoulBusRouteInfo.ed_station_nm!,
                maximum_dispatch_interval: seoulBusRouteInfo.term!,
                minimum_dispatch_interval: seoulBusRouteInfo.term!,
                start_station_first_time: parseTime(seoulBusRouteInfo.first_bus_tm),
                start_station_last_time: parseTime(seoulBusRouteInfo.last_bus_tm),
                start_station_low_first_time: parseTime(seoulBusRouteInfo.first_low_tm),
                start_station_low_last_time: parseTime(seoulBusRouteInfo.last_low_tm),
                end_station_first_time: null,
                end_station_last_time: null
            });
        }
    );
    console.log(`     [서울] 버스 노선 정보 가공 완료`);
}

const parseRouteType = (route_type: string): string => {
    switch(route_type) {
        case "3":
        case "4":
        case "5":
            return 'CITY';
        case "7":
        case "8":
            return 'INTERCITY';
        case "2":
            return 'VILLAGE';
        case "6":
            return 'WIDE_AREA';
        case "1":
            return 'AIRPORT';
        default:
            return 'OTHER';
    }
}

const parseTime = (time: string | null): string | null => {
    if (time?.trim()) {
        return dayjs(time.trim(), 'yyyyMMddHHmmss').format('HHmm');
    }
    return null;   
}

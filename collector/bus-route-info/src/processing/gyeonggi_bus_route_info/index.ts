import { OriginalGyeonggiBusRouteInfo } from "@prisma/client";
import { findManyOriginalGyeonggiBusRouteInfo, upsert } from "./repository";

export const processGyeonggiBusRouteInfo = async (): Promise<void> => {
    const gyeonggiBusRouteInfoList: OriginalGyeonggiBusRouteInfo[] = await findManyOriginalGyeonggiBusRouteInfo();
    gyeonggiBusRouteInfoList
        .filter((gyeonggiBusRouteInfo) => gyeonggiBusRouteInfo.route_name
            && gyeonggiBusRouteInfo.route_type_cd
            && gyeonggiBusRouteInfo.start_station_name
            && gyeonggiBusRouteInfo.end_station_name
            && gyeonggiBusRouteInfo.n_peek_alloc
            && gyeonggiBusRouteInfo.peek_alloc
        )
        .map((gyeonggiBusRouteInfo) => {
            upsert({
                route_id: gyeonggiBusRouteInfo.route_id,
                route_name: gyeonggiBusRouteInfo.route_name!,
                route_type: parseRouteType(gyeonggiBusRouteInfo.route_type_cd!),
                start_station_name: gyeonggiBusRouteInfo.start_station_name!,
                end_station_name: gyeonggiBusRouteInfo.end_station_name!,
                maximum_dispatch_interval: gyeonggiBusRouteInfo.n_peek_alloc!,
                minimum_dispatch_interval: gyeonggiBusRouteInfo.peek_alloc!,
                start_station_first_time: parseTime(gyeonggiBusRouteInfo.up_first_time),
                start_station_last_time: parseTime(gyeonggiBusRouteInfo.up_last_time),
                start_station_low_first_time: null,
                start_station_low_last_time: null,
                end_station_first_time: parseTime(gyeonggiBusRouteInfo.down_first_time),
                end_station_last_time: parseTime(gyeonggiBusRouteInfo.down_last_time),
            })
        }
    );
    console.log(`     [경기] 버스 노선 정보 가공 완료`);
}

const parseRouteType = (route_type_cd: string): string => {
    switch(route_type_cd) {
        case "11":
        case "12":
        case "13":
        case "15":
        case "16":
            return 'CITY';
        case "41":
        case "42":
        case "43":
            return 'INTERCITY';
        case "30":
            return 'VILLAGE';
        case "14":
            return 'WIDE_AREA';
        case "51":
        case "52":
        case "53":
            return 'AIRPORT';
        default:
            return 'OTHER';
    }
}

const parseTime = (time: string | null): string | null => {
    if (time?.trim()) {
        return time.trim().replace(':', '');
    }
    return null;
}

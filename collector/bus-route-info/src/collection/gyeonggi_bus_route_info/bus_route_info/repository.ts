import dayjs from "dayjs";
import { OriginalGyeonggiBusRouteInfo } from "@prisma/client";

import { prisma } from "../../../database";
import { GyeonggiBusRouteInfoData } from "./interface";

export const upsert = async (gyeonggiBusRouteInfoData: GyeonggiBusRouteInfoData): Promise<void> => {
    await prisma.originalGyeonggiBusRouteInfo.upsert({
        where: {
            route_id: gyeonggiBusRouteInfoData.routeId,
        },
        update: {
            route_name: gyeonggiBusRouteInfoData.routeName,
            route_type_cd: gyeonggiBusRouteInfoData.routeTypeCd,
            route_type_name: gyeonggiBusRouteInfoData.routeTypeName,
            start_station_id: gyeonggiBusRouteInfoData.startStationId,
            start_station_name: gyeonggiBusRouteInfoData.startStationName,
            start_mobile_no: gyeonggiBusRouteInfoData.startMobileNo,
            end_station_id: gyeonggiBusRouteInfoData.endStationId,
            end_station_name: gyeonggiBusRouteInfoData.endStationName,
            region_name: gyeonggiBusRouteInfoData.regionName,
            district_cd: gyeonggiBusRouteInfoData.districtCd,
            up_first_time: gyeonggiBusRouteInfoData.upFirstTime,
            up_last_time: gyeonggiBusRouteInfoData.upLastTime,
            down_first_time: gyeonggiBusRouteInfoData.downFirstTime,
            down_last_time: gyeonggiBusRouteInfoData.downLastTime,
            peek_alloc: gyeonggiBusRouteInfoData.peekAlloc,
            n_peek_alloc: gyeonggiBusRouteInfoData.nPeekAlloc,
            company_id: gyeonggiBusRouteInfoData.companyId,
            company_name: gyeonggiBusRouteInfoData.companyName,
            company_tel: gyeonggiBusRouteInfoData.companyTel,
            collection_date: dayjs().format('YYYY-MM-DD')
        },
        create: {
            route_id: gyeonggiBusRouteInfoData.routeId,
            route_name: gyeonggiBusRouteInfoData.routeName,
            route_type_cd: gyeonggiBusRouteInfoData.routeTypeCd,
            route_type_name: gyeonggiBusRouteInfoData.routeTypeName,
            start_station_id: gyeonggiBusRouteInfoData.startStationId,
            start_station_name: gyeonggiBusRouteInfoData.startStationName,
            start_mobile_no: gyeonggiBusRouteInfoData.startMobileNo,
            end_station_id: gyeonggiBusRouteInfoData.endStationId,
            end_station_name: gyeonggiBusRouteInfoData.endStationName,
            region_name: gyeonggiBusRouteInfoData.regionName,
            district_cd: gyeonggiBusRouteInfoData.districtCd,
            up_first_time: gyeonggiBusRouteInfoData.upFirstTime,
            up_last_time: gyeonggiBusRouteInfoData.upLastTime,
            down_first_time: gyeonggiBusRouteInfoData.downFirstTime,
            down_last_time: gyeonggiBusRouteInfoData.downLastTime,
            peek_alloc: gyeonggiBusRouteInfoData.peekAlloc,
            n_peek_alloc: gyeonggiBusRouteInfoData.nPeekAlloc,
            company_id: gyeonggiBusRouteInfoData.companyId,
            company_name: gyeonggiBusRouteInfoData.companyName,
            company_tel: gyeonggiBusRouteInfoData.companyTel,
            collection_date: dayjs().format('YYYY-MM-DD')
        }
    });
}

export const findMany = async (): Promise<OriginalGyeonggiBusRouteInfo[]> => {
    return await prisma.originalGyeonggiBusRouteInfo.findMany({ orderBy: { collection_date: 'desc' } });
}

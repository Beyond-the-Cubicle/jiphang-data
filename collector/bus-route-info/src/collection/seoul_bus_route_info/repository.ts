import dayjs from "dayjs";

import { prisma } from "../../database";
import { SeoulBusRouteInfoData } from "./interface";

export const saveAll = async (seoulBusRouteInfoDataList: SeoulBusRouteInfoData[]): Promise<void> => {
    await prisma.originalSeoulBusRouteInfo.createMany({
        data: seoulBusRouteInfoDataList.map((seoulBusRouteInfoData) => ({
            collection_date: dayjs().format('YYYY-MM-DD'),
            bus_route_id: seoulBusRouteInfoData.busRouteId,
            bus_route_nm: seoulBusRouteInfoData.busRouteNm,
            bus_route_abrv: seoulBusRouteInfoData.busRouteAbrv,
            length: seoulBusRouteInfoData.length,
            route_type: seoulBusRouteInfoData.routeType,
            st_station_nm: seoulBusRouteInfoData.stStationNm,
            ed_station_nm: seoulBusRouteInfoData.edStationNm,
            term: seoulBusRouteInfoData.term,
            last_bus_yn: seoulBusRouteInfoData.lastBusYn,
            first_bus_tm: seoulBusRouteInfoData.firstBusTm,
            last_bus_tm: seoulBusRouteInfoData.lastBusTm,
            first_low_tm: seoulBusRouteInfoData.firstLowTm,
            last_low_tm: seoulBusRouteInfoData.lastLowTm,
            corp_nm: seoulBusRouteInfoData.corpNm,
        })),
        skipDuplicates: true,
    });
}

export const deleteAllByCollectionDate = async (date: string): Promise<void> => {
    await prisma.originalSeoulBusRouteInfo.deleteMany({
        where: { collection_date: date }
    });
}

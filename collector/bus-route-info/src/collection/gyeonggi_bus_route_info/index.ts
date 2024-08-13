import dotenv from 'dotenv';

import { collectGyeonggiBusRoute, getGyeonggiBusRouteIdList } from "./bus_route";
import { collectGyeonggiBusRouteInfo, getLatestUpdatedGyeonggiBusRouteIdList } from "./bus_route_info";
import { constrainedMemory } from 'process';

dotenv.config();

const DATA_PORTAL_SERVICE_KEY_COUNT = process.env.DATA_PORTAL_API_KEY!.split(",").length;
const DATA_PORTAL_SERVICE_REQUEST_LIMIT = 900;
const DAILY_THROUGHPUT = DATA_PORTAL_SERVICE_KEY_COUNT * DATA_PORTAL_SERVICE_REQUEST_LIMIT;

export const collectGyeonggiBusRoutInfo = async (): Promise<void> => {
    await collectGyeonggiBusRoute();
    const targetGyeonggiBusRouteIdList = await getTargetGyeonggiBusRouteIdList();
    await collectGyeonggiBusRouteInfo(targetGyeonggiBusRouteIdList);
    console.log(`     [경기] 버스 노선 정보 수집 완료`);
}

const getTargetGyeonggiBusRouteIdList = async (): Promise<string[]> => {
    const gyeonggiBusRouteIdList: string[] = await getGyeonggiBusRouteIdList();
    if (gyeonggiBusRouteIdList.length < DAILY_THROUGHPUT) {
        return gyeonggiBusRouteIdList;
    }
    const latestUpdatedGyeonggiBusRouteIdList: string[] = await getLatestUpdatedGyeonggiBusRouteIdList();
    if (gyeonggiBusRouteIdList.length === latestUpdatedGyeonggiBusRouteIdList.length) {
        return latestUpdatedGyeonggiBusRouteIdList.splice(latestUpdatedGyeonggiBusRouteIdList.length - DAILY_THROUGHPUT);
    }
    if (latestUpdatedGyeonggiBusRouteIdList.length <= gyeonggiBusRouteIdList.length - DAILY_THROUGHPUT) {
        const targetGyeonggiBusRouteIdList: string[] = gyeonggiBusRouteIdList.filter(gyeonggiBusRouteId => !latestUpdatedGyeonggiBusRouteIdList.includes(gyeonggiBusRouteId));
        return targetGyeonggiBusRouteIdList.length > DAILY_THROUGHPUT ? targetGyeonggiBusRouteIdList.splice(0, DAILY_THROUGHPUT) : targetGyeonggiBusRouteIdList;
    }
    if (latestUpdatedGyeonggiBusRouteIdList.length > gyeonggiBusRouteIdList.length - DAILY_THROUGHPUT) {
        const newGyeonggiBusRouteIdList: string[] = gyeonggiBusRouteIdList.filter(gyeonggiBusRouteId => !latestUpdatedGyeonggiBusRouteIdList.includes(gyeonggiBusRouteId)); 
        latestUpdatedGyeonggiBusRouteIdList.splice(latestUpdatedGyeonggiBusRouteIdList.length - DAILY_THROUGHPUT + newGyeonggiBusRouteIdList.length);
        return gyeonggiBusRouteIdList.filter(gyeonggiBusRouteId => !latestUpdatedGyeonggiBusRouteIdList.includes(gyeonggiBusRouteId));
    }
    console.warn('     [경기] 버스 노선 정보 수집 대상 식별 실패')
    return [];
}

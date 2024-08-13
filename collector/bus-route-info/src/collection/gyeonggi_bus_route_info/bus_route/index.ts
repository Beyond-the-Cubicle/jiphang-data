
import { OriginalGyeonggiBusRoute } from '@prisma/client';

import { fetchGyeonggiBusRoute } from "./request";
import { GyeonggiBusRouteResponse, GyeonggiBusRouteData, GyeonggiBusRouteErrorResponse } from "./interface";
import { upsert, findMany } from './repository';

const DATA_DREAM_SERVICE_KEY: string = process.env.DATA_DREAM_API_KEY!;
const DEFAULT_START_PAGE_INDEX: number = 1;
const DEFAULT_PAGE_MAX_SIZE: number = 1000;

export const collectGyeonggiBusRoute = async (): Promise<GyeonggiBusRouteData[]> => {
    const gyeonggiBusRouteDataList: GyeonggiBusRouteData[] = await getGyeonggiBusRouteDataList();
    gyeonggiBusRouteDataList.map((gyeonggiBusRouteData) => upsert(gyeonggiBusRouteData));
    return gyeonggiBusRouteDataList;
}

export const getGyeonggiBusRouteIdList = async (): Promise<string[]> => {
    const gyeonggiBusRouteDataList: OriginalGyeonggiBusRoute[] = await findMany();
    return gyeonggiBusRouteDataList.map((gyeonggiBusRouteData) => gyeonggiBusRouteData.route_id);
}

const getGyeonggiBusRouteDataList = async (): Promise<GyeonggiBusRouteData[]> => {
    const gyeonggiBusRouteDataList: GyeonggiBusRouteData[] = [];
    const gyeonggiBusRouteTotalPage: number = await getGyeonggiBusRouteTotalPage();

    for (let page = DEFAULT_START_PAGE_INDEX; page <= gyeonggiBusRouteTotalPage; page++) {
        const response = await getGyeonggiBusRoute(page, DEFAULT_PAGE_MAX_SIZE);
        response.TBBMSROUTEM.row.map(gyeonggiBusRouteData => gyeonggiBusRouteDataList.push(gyeonggiBusRouteData));
    }

    return gyeonggiBusRouteDataList;
}

const getGyeonggiBusRouteTotalPage = async (): Promise<number> => {
    const response: GyeonggiBusRouteResponse = await getGyeonggiBusRoute(DEFAULT_START_PAGE_INDEX, DEFAULT_PAGE_MAX_SIZE);
    return Math.floor((response as GyeonggiBusRouteResponse).TBBMSROUTEM.head.list_total_count / DEFAULT_PAGE_MAX_SIZE) + 1;
}

const getGyeonggiBusRoute = async (pIndex: number, pSize: number): Promise<GyeonggiBusRouteResponse> => {
    const response: GyeonggiBusRouteResponse | GyeonggiBusRouteErrorResponse = await fetchGyeonggiBusRoute(DATA_DREAM_SERVICE_KEY, pIndex, pSize);
    if (!isResponseSuccessful(response)) {
        throw new Error(`     [경기] 버스 노선 요청 실패: ${(response as GyeonggiBusRouteErrorResponse).RESULT.MESSAGE}`);
    }
    return response as GyeonggiBusRouteResponse;
}

const isResponseSuccessful = (response: GyeonggiBusRouteResponse | GyeonggiBusRouteErrorResponse): boolean => {
    if (gyeonggiBusRouteResponseTypeAssertion(response)) {
        return response.TBBMSROUTEM.head.RESULT.CODE === 'INFO-000';
    }

    if (gyeonggiBusRouteErrorResponseTypeAssertion(response)) {
        return false;
    } 

    return false;
}

const gyeonggiBusRouteResponseTypeAssertion = (arg: any): arg is GyeonggiBusRouteResponse => {
    return arg.TBBMSROUTEM !== undefined;
}

const gyeonggiBusRouteErrorResponseTypeAssertion = (arg: any): arg is GyeonggiBusRouteErrorResponse => {
    return arg.RESULT !== undefined;
}

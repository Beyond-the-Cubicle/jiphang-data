import { OriginalGyeonggiBusRouteInfo } from "@prisma/client";

import { delay } from "../../../util";
import { fetchGyeonggiBusRouteInfo } from "./request";
import { GyeonggiBusRouteInfoResponse, GyeonggiBusRouteInfoData, GyeonggiBusRouteInfoErrorResponse } from "./interface";
import { upsert, findMany } from "./repository";

const DATA_PORTAL_REQUEST_BATCH_SIZE: number = 1;
const DATA_PORTAL_REQUEST_BATCH_SLEEP: number = 100;
const DATA_PORTAL_SERVICE_KEY_LIST: string[] = process.env.DATA_PORTAL_API_KEY!.split(",");

export const collectGyeonggiBusRouteInfo = async (gyeonggiBusRouteIdList: string[]): Promise<void> => {
    const pLimit = (await import('p-limit')).default;
    const limit = pLimit(DATA_PORTAL_REQUEST_BATCH_SIZE);

    for (let i = 0; i < gyeonggiBusRouteIdList.length; i += DATA_PORTAL_REQUEST_BATCH_SIZE) {
        const batch = gyeonggiBusRouteIdList.slice(i, i + DATA_PORTAL_REQUEST_BATCH_SIZE);
        await Promise.all(batch.map(routeId => limit(() => getGyeonggiBusRouteInfoData(routeId)
                .then((gyeonggiBusRouteInfoData) => upsert(gyeonggiBusRouteInfoData))
                .catch((error) => console.error(error))
            )));
        if (i + DATA_PORTAL_REQUEST_BATCH_SIZE < gyeonggiBusRouteIdList.length) {
            await delay(DATA_PORTAL_REQUEST_BATCH_SLEEP);
        }
    }
}

export const getLatestUpdatedGyeonggiBusRouteIdList = async (): Promise<string[]> => {
    const gyeonggiBusRouteInfoDataList: OriginalGyeonggiBusRouteInfo[] = await findMany();
    return gyeonggiBusRouteInfoDataList.map((gyeonggiBusRouteInfoData) => gyeonggiBusRouteInfoData.route_id);
}

const getGyeonggiBusRouteInfoData = async (routeId: string): Promise<GyeonggiBusRouteInfoData> => {
    const response = await fetchGyeonggiBusRouteInfo(getDataPortalKey(), routeId);
    if (gyeonggiBusRouteInfoResponseTypeAssertion(response)) {
        if (response.response.msgHeader) {
            switch(response.response.msgHeader.resultCode) {
                case '0':
                    return response.response.msgBody!.busRouteInfoItem;
                case '4':
                    return { routeId: routeId };
                default:
                    throw new Error(`     [경기] 버스 노선 정보 요청 실패: ${response.response.msgHeader.resultMessage}`);
            }
        } else {
            console.warn(`     [경기] 버스 노선 정보 요청 재시도: msgHeader 빈값 응답`);
            return await getGyeonggiBusRouteInfoData(routeId);
        }
        
    }
    if (gyeonggiBusRouteInfoErrorResponseTypeAssertion(response)) {
        if(!DATA_PORTAL_SERVICE_KEY_LIST[0]) {
            console.warn('     [경기] 버스 노선 정보 수집 요청 허용 범위 초과')
            process.exit(1);
        }
        switch(response.OpenAPI_ServiceResponse.cmmMsgHeader.returnReasonCode) {
            case '22':
                DATA_PORTAL_SERVICE_KEY_LIST.shift();
                return await getGyeonggiBusRouteInfoData(routeId);
            default:
                throw new Error(`     [경기] 버스 노선 정보 요청 실패: ${response.OpenAPI_ServiceResponse.cmmMsgHeader.returnAuthMsg}`);
        }
    }
    throw new Error(`     [경기] 버스 노선 정보 요청 실패`);
}

const gyeonggiBusRouteInfoResponseTypeAssertion = (arg: any): arg is GyeonggiBusRouteInfoResponse => {
    return arg.response !== undefined;
}

const gyeonggiBusRouteInfoErrorResponseTypeAssertion = (arg: any): arg is GyeonggiBusRouteInfoErrorResponse => {
    return arg.OpenAPI_ServiceResponse !== undefined;
}

const getDataPortalKey = (): string => {
    return DATA_PORTAL_SERVICE_KEY_LIST[0];
}

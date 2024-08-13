import dayjs from "dayjs";

import { fetchSeoulBusRouteInfo } from "./request";
import { SeoulBusRouteInfoResponse, SeoulBusRouteInfoData } from "./interface";
import { saveAll, deleteAllByCollectionDate } from "./repository";

const DATA_PORTAL_SERVICE_KEY = process.env.DATA_PORTAL_API_KEY!.split(",")[0];

export const collectSeoulBusRoutInfo = async (): Promise<void> => {
    try {
        const seoulBusRouteInfoDataList: SeoulBusRouteInfoData[] = await getSeoulBusRouteInfoDataList();
        await saveAll(seoulBusRouteInfoDataList);
        await applyRollingPolicy();
        console.log(`     [서울] 버스 노선 정보 수집 완료`);
    } catch (error) {
        console.error(error.message);
    }
}

const getSeoulBusRouteInfoDataList = async (): Promise<SeoulBusRouteInfoData[]> => {
    const response = await fetchSeoulBusRouteInfo(DATA_PORTAL_SERVICE_KEY);
    if (!isResponseSuccessful(response)) {
        throw new Error(`     [서울] 버스 노선 정보 요청 실패: ${response.ServiceResult.msgHeader.headerMsg}`);
    }
    return response.ServiceResult.msgBody!.itemList;
}

const isResponseSuccessful = (response: SeoulBusRouteInfoResponse): boolean => {
    return response.ServiceResult.msgHeader.headerCd === '0';
}

const applyRollingPolicy = async (): Promise<void> => {
    await deleteAllByCollectionDate(dayjs().subtract(2, 'days').format('YYYY-MM-DD'));
}

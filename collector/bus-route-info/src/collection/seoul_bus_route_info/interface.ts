export interface SeoulBusRouteInfoResponse {
    ServiceResult: {
        comMsgHeader: undefined;
        msgHeader: {
            headerCd: string;
            headerMsg: string;
            itemCount: string;
        },
        msgBody: {
            itemList: SeoulBusRouteInfoData[];
        }
    }
}

export interface SeoulBusRouteInfoData {
    busRouteAbrv: string;
    busRouteId: string;
    busRouteNm: string;
    corpNm: string;
    edStationNm: string;
    firstBusTm: string;
    firstLowTm: string;
    lastBusTm: string;
    lastBusYn: string;
    lastLowTm: string;
    length: string;
    routeType: string;
    stStationNm: string;
    term: string;
}

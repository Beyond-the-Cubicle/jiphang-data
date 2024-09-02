export interface GyeonggiBusRouteInfoResponse {
    response: {
        comMsgHeader?: {
            errMsg: string
            returnAuthMsg: string
            returnReasonCode: string
        } | string;
        msgHeader: {
            queryTime: string;
            resultCode: string;
            resultMessage: string;
        },
        msgBody?: {
            busRouteInfoItem: GyeonggiBusRouteInfoData;
        }
    }
}

export interface GyeonggiBusRouteInfoData {
    routeId: string;
    routeName?: string;
    routeTypeCd?: string;
    routeTypeName?: string;
    startStationId?: string;
    startStationName?: string;
    startMobileNo?: string;
    endStationId?: string;
    endStationName?: string;
    regionName?: string;
    districtCd?: string;
    upFirstTime?: string;
    upLastTime?: string;
    downFirstTime?: string;
    downLastTime?: string;
    peekAlloc?: string;
    nPeekAlloc?: string;
    companyId?: string;
    companyName?: string;
    companyTel?: string;
}

export interface GyeonggiBusRouteInfoErrorResponse {
    OpenAPI_ServiceResponse: {
        cmmMsgHeader: {
            errMsg: string
            returnAuthMsg: string
            returnReasonCode: string
        };
    }
}

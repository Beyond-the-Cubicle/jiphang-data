import { requestAndParseXml } from "../../request";
import { GyeonggiBusRouteInfoResponse, GyeonggiBusRouteInfoErrorResponse } from "./interface";

const BASE_URL = "http://apis.data.go.kr/6410000/busrouteservice/getBusRouteInfoItem";

export const fetchGyeonggiBusRouteInfo = async (serviceKey: string, routeId: string): Promise<GyeonggiBusRouteInfoResponse | GyeonggiBusRouteInfoErrorResponse> => {
    const url: string = `${BASE_URL}?serviceKey=${serviceKey}&routeId=${routeId}`;
    return requestAndParseXml<GyeonggiBusRouteInfoResponse | GyeonggiBusRouteInfoErrorResponse>(url);
}

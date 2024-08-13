import { requestAndParseXml } from "../request";
import { SeoulBusRouteInfoResponse } from "./interface";

const BASE_URL = "http://ws.bus.go.kr/api/rest/busRouteInfo/getBusRouteList";

export const fetchSeoulBusRouteInfo = async (serviceKey: string): Promise<SeoulBusRouteInfoResponse> => {
    const url = `${BASE_URL}?serviceKey=${serviceKey}`;
    return requestAndParseXml<SeoulBusRouteInfoResponse>(url);
}

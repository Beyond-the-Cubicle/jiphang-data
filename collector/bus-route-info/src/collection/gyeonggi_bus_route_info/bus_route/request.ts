import { requestAndParseXml } from "../../request";
import { GyeonggiBusRouteResponse, GyeonggiBusRouteErrorResponse } from "./interface";

const BASE_URL = "https://openapi.gg.go.kr/TBBMSROUTEM";

export const fetchGyeonggiBusRoute = async (serviceKey: string, pageIndex: number, pageSize: number): Promise<GyeonggiBusRouteResponse | GyeonggiBusRouteErrorResponse> => {
    const url: string = `${BASE_URL}?KEY=${serviceKey}&Type=xml&pIndex=${pageIndex}&pSize=${pageSize}`;
    return requestAndParseXml<GyeonggiBusRouteResponse | GyeonggiBusRouteErrorResponse>(url);
};

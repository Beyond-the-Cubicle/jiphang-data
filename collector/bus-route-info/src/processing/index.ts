import { processGyeonggiBusRouteInfo } from "./gyeonggi_bus_route_info";
import { processSeoulBusRouteInfo } from "./seoul_bus_route_info";

export const processing = async (): Promise<void> => {
    console.log('  [PROCESSING] :: 버스 노선 정보 가공')
    // 버스 노선 우선 순위 기준: 서울
    await processSeoulBusRouteInfo();
    await processGyeonggiBusRouteInfo();
}

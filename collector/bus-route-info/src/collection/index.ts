import { collectGyeonggiBusRoutInfo } from "./gyeonggi_bus_route_info";
import { collectSeoulBusRoutInfo } from "./seoul_bus_route_info";

export const collection = async (): Promise<void> => {
    console.log('  [COLLECTION] :: 버스 노선 정보 수집')
    await Promise.all([
        collectGyeonggiBusRoutInfo(),
        collectSeoulBusRoutInfo()
    ]).catch(error => console.error(error));
}

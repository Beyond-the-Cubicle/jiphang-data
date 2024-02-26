-- chulgeun_gil_planner.seoul_bus_link definition

CREATE TABLE `seoul_bus_link` (
  `routeId` bigint DEFAULT NULL COMMENT '노선 ID',
  `stationId` bigint DEFAULT NULL COMMENT '정류장 ID',
  `stationDistanceMeter` int DEFAULT NULL COMMENT '링크 구간거리(m)',
  `stationOrder` int DEFAULT NULL COMMENT '정류장 순서 번호'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

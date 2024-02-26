-- chulgeun_gil_planner.seoul_bus_link definition

CREATE TABLE `seoul_bus_link` (
  `routeId` bigint DEFAULT NULL COMMENT '노선 ID',
  `stationId` bigint DEFAULT NULL COMMENT '정류장 ID',
  `stationDistanceMeter` int DEFAULT NULL COMMENT '링크 구간거리(m)',
  `stationOrder` int DEFAULT NULL COMMENT '정류장 순서 번호'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- chulgeun_gil_planner.gyunggi_bus_link definition

CREATE TABLE `gyunggi_bus_link` (
  `routeId` bigint DEFAULT NULL COMMENT '노선 ID',
  `stationOrder` int DEFAULT NULL COMMENT '정류장 순서 번호',
  `stationId` bigint DEFAULT NULL COMMENT '정류장 ID',
  `gisDistance` int DEFAULT NULL COMMENT 'GIS 거리',
  `accumulatedDistance` int DEFAULT NULL COMMENT '누적거리',
  `realDistance` int DEFAULT NULL COMMENT '실제거리',
  `decidedDistance` int DEFAULT NULL COMMENT '확정거리',
  `progressCode` varchar(20) DEFAULT NULL COMMENT '진행구분코드',
  `registeredBy` varchar(100) DEFAULT NULL COMMENT '등록아이디',
  `registeredAt` varchar(50) DEFAULT NULL COMMENT '등록일자',
  `useDivision` varchar(100) DEFAULT NULL COMMENT '사용구분',
  `regionalLineYn` varchar(30) DEFAULT NULL COMMENT '벽지노선 유무',
  `progressCodeName` varchar(50) DEFAULT NULL COMMENT '진행구분코드명',
  `useDivisionName` varchar(50) DEFAULT NULL COMMENT '벽지노선 유무명'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

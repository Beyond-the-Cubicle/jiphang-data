CREATE DATABASE `chulgeun_gil_planner` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

-- chulgeun_gil_planner.gyunggi_bus_station definition

CREATE TABLE `gyunggi_bus_station` (
  `stationId` varchar(100) NOT NULL COMMENT '정류장 ID',
  `stationName` varchar(200) NOT NULL COMMENT '정류장 이름',
  `coordinateX` float DEFAULT NULL COMMENT 'X좌표',
  `coordinateY` float DEFAULT NULL COMMENT 'Y좌표',
  `gpsCoordinateX` float DEFAULT NULL COMMENT 'GPS X좌표',
  `gpsCoordinateY` float DEFAULT NULL COMMENT 'GPS Y좌표',
  `rinkId` varchar(100) DEFAULT NULL COMMENT '링크아이디',
  `stationType` varchar(100) DEFAULT NULL COMMENT '정류장 유형',
  `transferStationExtNo` varchar(20) DEFAULT NULL COMMENT '환승 정류장 유무 ex. C, N',
  `medianBusLaneYn` varchar(20) DEFAULT NULL COMMENT '중앙차로 여부',
  `stationEnglishName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '정류장 영문명',
  `arsId` varchar(100) DEFAULT NULL COMMENT 'ARS ID',
  `institutionCode` varchar(100) DEFAULT NULL COMMENT '기관 코드',
  `dataDisplayYn` varchar(20) DEFAULT NULL COMMENT '데이터 표출 유무(Y/N)',
  `registeredBy` varchar(100) DEFAULT NULL COMMENT '등록 주체 아이디',
  `registeredAt` varchar(100) DEFAULT NULL COMMENT '등록 일시 YYYYMMDDHHmmss',
  `memo` varchar(100) DEFAULT NULL COMMENT '비고',
  `signPostType` varchar(100) DEFAULT NULL COMMENT '표지판 유형',
  `dongCode` varchar(100) DEFAULT NULL COMMENT '행정동 코드',
  `regionCode` varchar(100) DEFAULT NULL COMMENT '권역 코드',
  `useYn` varchar(100) DEFAULT NULL COMMENT '사용구분(Y/N)',
  `stationChineseName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '정류장 중국어 이름',
  `stationJapaneseName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '정류장 일본어 이름',
  `stationVietnamName` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '정류장 베트남어 이름',
  `drtYn` varchar(20) DEFAULT NULL COMMENT 'DRT 유무',
  `stationTypeName` varchar(100) DEFAULT NULL COMMENT '정류장 유형 이름(ex. 미지정, 시내)',
  `trensferStationTypeName` varchar(100) DEFAULT NULL COMMENT '환승유무(ex. 일반, 환승)',
  `signPostTypeName` varchar(100) DEFAULT NULL COMMENT '표지판 유형 이름(ex. 표지판 없음)',
  PRIMARY KEY (`stationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- chulgeun_gil_planner.seoul_bus_station definition

CREATE TABLE `seoul_bus_station` (
  `stationId` varchar(100) NOT NULL COMMENT '정류장 ID',
  `stationName` varchar(200) NOT NULL COMMENT '정류장 이름',
  `stationType` varchar(100) DEFAULT NULL COMMENT '정류장 유형(ex. 일반차로, 중앙차로...)',
  `arsId` varchar(20) DEFAULT NULL COMMENT '정류장 번호 - ARS ID',
  `coordinateX` float DEFAULT NULL COMMENT '경도',
  `coordinateY` float DEFAULT NULL COMMENT '위도',
  `busArrivalInfoGuideInstallYn` varchar(20) DEFAULT NULL COMMENT '버스도착정보안내기 설치 여부',
  PRIMARY KEY (`stationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- chulgeun_gil_planner.standard_bus_station definition

CREATE TABLE `standard_bus_station` (
  `stationId` varchar(100) NOT NULL COMMENT '정류소 ID',
  `stationName` varchar(200) NOT NULL COMMENT '정류소 이름',
  `arsId` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '정류소 번호',
  `latitude` double NOT NULL COMMENT '정류소 위도',
  `longitude` double NOT NULL COMMENT '정류소 경도',
  PRIMARY KEY (`stationId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

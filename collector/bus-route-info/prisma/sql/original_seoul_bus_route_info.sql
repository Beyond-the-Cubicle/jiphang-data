create table original_seoul_bus_route_info (

    collection_date                 varchar(191)    not null    comment '수집 날짜',
    bus_route_id                    varchar(191)    not null    comment '노선 ID',
    bus_route_nm                    varchar(191)    null        comment '노선명(DB관리용)',
    bus_route_abrv                  varchar(191)    null        comment '노선명(안내용 - 마을버스 제외)',
    length                          varchar(191)    null        comment '노선 길이 (Km)',
    route_type                      varchar(191)    null        comment '노선 유형 (1:공항, 2:마을, 3:간선, 4:지선, 5:순환, 6:광역, 7:인천, 8:경기, 9:폐지, 0:공용)',
    st_station_nm                   varchar(191)    null        comment '기점',
    ed_station_nm                   varchar(191)    null        comment '종점',
    term                            varchar(191)    null        comment '배차간격 (분)',
    last_bus_yn                     varchar(191)    null        comment '막차 운행 여부',
    first_bus_tm                    varchar(191)    null        comment '금일 첫차 시간',
    last_bus_tm                     varchar(191)    null        comment '금일 막차 시간',
    first_low_tm                    varchar(191)    null        comment '금일 저상 첫차 시간',
    last_low_tm                     varchar(191)    null        comment '금일 저상 막차 시간',
    corp_nm                         varchar(191)    null        comment '운수사명',

    primary key(collection_date, bus_route_id)
);

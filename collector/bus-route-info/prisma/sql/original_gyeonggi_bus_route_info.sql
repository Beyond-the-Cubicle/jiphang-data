create table original_gyeonggi_bus_route_info (

    route_id                        varchar(191)    not null    comment '노선 아이디',
    route_name                      varchar(191)    null        comment '노선 번호',
    route_type_cd                   varchar(191)    null        comment '노선 유형',
    route_type_name                 varchar(191)    null        comment '노선 유형명',
    start_station_id                varchar(191)    null        comment '기점 정류소 아이디',
    start_station_name              varchar(191)    null        comment '기점 정류소명',
    start_mobile_no                 varchar(191)    null        comment '기점 정류소번호',
    end_station_id                  varchar(191)    null        comment '종점 정류소 아이디',
    end_station_name                varchar(191)    null        comment '종점 정류소명',
    region_name                     varchar(191)    null        comment '지역 명',
    district_cd                     varchar(191)    null        comment '관할 지역',
    up_first_time                   varchar(191)    null        comment '평일 기점 첫차시간',
    up_last_time                    varchar(191)    null        comment '평일 기점 막차시간',
    down_first_time                 varchar(191)    null        comment '평일 종점 첫차시간',
    down_last_time                  varchar(191)    null        comment '평일 종점 막차시간',
    peek_alloc                      varchar(191)    null        comment '평일 최소 배차시간',
    n_peek_alloc                    varchar(191)    null        comment '평일 최대 배차시간',
    company_id                      varchar(191)    null        comment '운수 업체 아이디',
    company_name                    varchar(191)    null        comment '운수 업체명',
    company_tel                     varchar(191)    null        comment '운수 업체 전화번호'
    collection_date                 varchar(191)    not null    comment '수집 날짜',

    primary key(route_id)
);

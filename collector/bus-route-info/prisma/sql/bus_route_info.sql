create table bus_route_info (

    route_id                        varchar(191)    not null    comment '노선 아이디',
    route_name                      varchar(191)    not null    comment '노선 번호',
    route_type                      varchar(191)    not null    comment '노선 유형',
    start_station_name              varchar(191)    not null    comment '기점정류소명',
    end_station_name                varchar(191)    not null    comment '종점정류소명',
    maximum_dispatch_interval       varchar(191)    not null    comment '평일 최대 배차시간',
    minimum_dispatch_interval       varchar(191)    not null    comment '평일 최소 배차시간',
    start_station_first_time        varchar(191)    null        comment '기점 첫차 시간',
    start_station_last_time         varchar(191)    null        comment '기점 막차 시간',
    start_station_low_first_time    varchar(191)    null        comment '기점 저상 첫차 시간',
    start_station_low_last_time     varchar(191)    null        comment '기점 저상 막차 시간',
    end_station_first_time          varchar(191)    null        comment '종점 첫차 시간',
    end_station_last_time           varchar(191)    null        comment '종점 막차 시간',

    primary key(route_id)
);

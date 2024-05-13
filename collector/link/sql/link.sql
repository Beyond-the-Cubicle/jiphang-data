create table link
(
    route_id         bigint not null,
    start_station_id bigint not null,
    end_station_id   bigint not null,
    trip_time        int    not null,
    trip_distance    int    not null,
    station_order    int    not null,
    primary key (route_id, start_station_id, end_station_id)
);


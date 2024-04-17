create table seoul_link
(
    route_id       varchar(191) not null,
    sttn_id        varchar(191) not null,
    sttn_dstnc_mtr int          not null,
    sttn_ordr      int          not null,
    primary key (route_id, sttn_id)
);


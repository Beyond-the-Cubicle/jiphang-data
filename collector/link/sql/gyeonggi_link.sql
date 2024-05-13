create table gyeonggi_link
(
    route_id                    varchar(191) not null,
    sttn_ordr                   int          not null,
    sttn_id                     varchar(191) not null,
    gis_dstn                    int          null,
    accmlt_dstn                 int          null,
    real_dstn                   int          null,
    dcsn_dstn                   int          null,
    progrs_div_cd               varchar(191) null,
    use_div                     varchar(191) null,
    unwel_hno_statn_route_extno varchar(191) null,
    progrs_div_cd_nm            varchar(191) null,
    use_div_nm                  varchar(191) null,
    primary key (route_id, sttn_id)
);


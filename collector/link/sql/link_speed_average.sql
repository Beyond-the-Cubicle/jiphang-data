create table link_speed_average
(
    year        int          not null,
    location    varchar(191) not null,
    weekdayKmh  double       not null,
    weekdayMs   double       not null,
    saturdayKmh double       not null,
    saturdayMs  double       not null,
    sundayKmh   double       not null,
    sundayMs    double       not null,
    primary key (year, location)
);


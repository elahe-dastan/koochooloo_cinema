CREATE TABLE registration
(
    username        varchar(255) primary key,
    password        varchar(255),
    firs_name       varchar(255),
    last_name       varchar(255),
    email           varchar(255),
    phone           varchar(13),
    national_number varchar(10)
)
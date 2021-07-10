CREATE TABLE registration
(
    username        varchar(255) primary key,
    password        varchar(255),
    firs_name       varchar(255),
    last_name       varchar(255),
    email           varchar(255),
    phone           varchar(13),
    national_number varchar(10),
    special_at      timestamp,
    credit          int default 0,
    score           int default 0,
    CONSTRAINT CHK_Password CHECK ( LENGTH(password) >= 8 AND  password LIKE '%[0-9]%' AND password LIKE '%[A-Z]%' AND password LIKE '%[a-z]%')
);

CREATE UNIQUE INDEX uidx_email ON registration (email);
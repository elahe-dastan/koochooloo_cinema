CREATE TABLE IF NOT EXISTS wallet
(
    username   varchar(255) primary key,
    credit     int default 0
)
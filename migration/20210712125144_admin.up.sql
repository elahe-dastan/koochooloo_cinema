CREATE TABLE IF NOT EXISTS admin
(
    username varchar(255) primary key,
    password varchar(255)
);

INSERT INTO admin VALUES ('admin', 'admin');
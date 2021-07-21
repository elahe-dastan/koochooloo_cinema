CREATE TABLE IF NOT EXISTS users
(
  username        varchar(255) primary key,
  password        varchar(255),
  first_name      varchar(255),
  last_name       varchar(255),
  email           varchar(255),
  phone           varchar(13),
  national_number varchar(10),
  special_till    date,
  score           int default 0,
  CONSTRAINT CHK_Password CHECK ( LENGTH(password) >= 8 AND  password LIKE '%[0-9]%' AND password LIKE '%[A-Z]%' AND password LIKE '%[a-z]%')
);

CREATE UNIQUE INDEX IF NOT EXISTS uidx_email ON users(email);

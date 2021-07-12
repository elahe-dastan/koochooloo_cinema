CREATE TABLE IF NOT EXISTS film
(
    id              int primary key auto_increment,
    file            varchar(255),
    name            varchar(255),
    production_year int,
    explanation     varchar(255),
    view            int default 0,
    price           int DEFAULT 0
)
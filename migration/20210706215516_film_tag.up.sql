CREATE TABLE IF NOT EXISTS film_tag
(
    film  int,
    tag varchar(255),
    primary key (film, tag),
    FOREIGN KEY (film) REFERENCES film (id)
)
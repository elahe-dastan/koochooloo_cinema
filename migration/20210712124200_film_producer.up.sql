CREATE TABLE IF NOT EXISTS film_producer
(
    film     int,
    producer varchar(255),
    primary key (film, producer),
    FOREIGN KEY (film) REFERENCES film (id)
)
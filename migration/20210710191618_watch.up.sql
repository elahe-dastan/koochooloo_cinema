CREATE TABLE IF NOT EXISTS watch (
    film_id int NOT NULL,
    user_id int NOT NULL,
    PRIMARY KEY (film_id, user_id),
    FOREIGN KEY (film_id) REFERENCES film(id),
    FOREIGN KEY (user_id) REFERENCES registeratioin(id)
)


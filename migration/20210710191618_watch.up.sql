CREATE TABLE IF NOT EXISTS watch
(
    film int,
    username varchar(255),
    number int,
    PRIMARY KEY (film, username),
    FOREIGN KEY (film) REFERENCES film (id),
    FOREIGN KEY (username) REFERENCES users (username)
)


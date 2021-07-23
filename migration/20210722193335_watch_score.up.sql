CREATE TABLE IF NOT EXISTS watch_score
(
    film int,
    username varchar(255),
    number int default 0,
    PRIMARY KEY (film, username),
    FOREIGN KEY (film) REFERENCES film (id),
    FOREIGN KEY (username) REFERENCES users (username)
);

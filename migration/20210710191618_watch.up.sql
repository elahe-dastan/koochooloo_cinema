CREATE TABLE IF NOT EXISTS watch
(
    film int,
    user varchar(255),
    number int,
    PRIMARY KEY (film, user),
    FOREIGN KEY (film) REFERENCES film (id),
    FOREIGN KEY (user) REFERENCES user (username)
)


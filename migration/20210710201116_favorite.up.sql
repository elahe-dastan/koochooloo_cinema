CREATE TABLE IF NOT EXISTS favorite
(
    user  varchar(255),
    film  int,
    album VARCHAR(255),
    primary key (user, film, album),
    FOREIGN KEY (user) REFERENCES user (username),
    FOREIGN KEY (film) REFERENCES film (id)
);


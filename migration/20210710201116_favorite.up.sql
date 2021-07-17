CREATE TABLE IF NOT EXISTS favorite
(
    username  varchar(255),
    film  int,
    album VARCHAR(255),
    primary key (username, film, album),
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (film) REFERENCES film (id)
);


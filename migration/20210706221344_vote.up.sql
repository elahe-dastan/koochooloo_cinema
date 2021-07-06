CREATE TABLE vote
(
    user  int,
    film  int,
    score int,
    comment varchar(255),
    primary key (user, film),
    FOREIGN KEY (user) REFERENCES registeration (id),
    FOREIGN KEY (film) REFERENCES film (id)
)

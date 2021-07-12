CREATE TABLE IF NOT EXISTS vote
(
    user    varchar(255),
    film    int,
    score   int,
    comment varchar(255),
    primary key (user, film),
    FOREIGN KEY (user) REFERENCES user (username),
    FOREIGN KEY (film) REFERENCES film (id)
);

CREATE TRIGGER watch_before_vote
    BEFORE
        INSERT
    ON vote
    for each row
    IF Not EXISTS(SELECT *
                  FROM watch
                  WHERE film_id = film
                    AND user_id = user)
BEGIN
    RAISERROR
    ("You havn't watche this film")
ROLLBACK
    END


#     for each row is wrong
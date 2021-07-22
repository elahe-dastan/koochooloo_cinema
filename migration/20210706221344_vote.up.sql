CREATE TABLE IF NOT EXISTS vote
(
    username    varchar(255),
    film    int,
    score   int,
    comment varchar(255),
    primary key (username, film),
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (film) REFERENCES film (id)
);

CREATE OR REPLACE FUNCTION watch_before_vote() RETURNS trigger as
$$
  BEGIN
    IF Not EXISTS(SELECT *
                FROM watch
                WHERE film = NEW.film
                  AND username = NEW.username) THEN
      RETURN NULL;
    END IF;

    UPDATE film set score = (score * ( select count(*) from vote where film = NEW.film ) + NEW.score) / ( select count(*) from vote where film = NEW.film ) + 1 where id = file;

    RETURN NEW;
  END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER vote_create BEFORE INSERT on vote FOR EACH ROW EXECUTE PROCEDURE watch_before_vote();

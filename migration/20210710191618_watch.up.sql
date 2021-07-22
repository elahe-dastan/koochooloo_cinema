CREATE TABLE IF NOT EXISTS watch
(
    film int,
    username varchar(255),
    number int default 0,
    PRIMARY KEY (film, username),
    FOREIGN KEY (film) REFERENCES film (id),
    FOREIGN KEY (username) REFERENCES users (username)
)


CREATE OR REPLACE FUNCTION pay_for_watch() RETURNS trigger as
$$
  BEGIN
    IF (select count(*) from user,film where film.id = NEW.film AND user.credit >= film.price AND username = NEW.username) THEN
      UPDATE film SET view = view + 1 WHERE id = NEW.film;
      UPDATE user SET credit = credit - ( SELECT price FROM film WHERE id = NEW.film ) WHERE username = NEW.username;
      RETURN NEW;
    END IF;
      RETURN NULL;
  END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER watch_create BEFORE INSERT on watch FOR EACH ROW EXECUTE PROCEDURE pay_for_watch();

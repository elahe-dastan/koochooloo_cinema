CREATE TABLE IF NOT EXISTS favorite
(
    username  varchar(255),
    film  int,
    album VARCHAR(255),
    primary key (username, film, album),
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (film) REFERENCES film (id)
);

CREATE OR REPLACE FUNCTION special_user_validation() RETURNS trigger as
$$
BEGIN
  IF (select count(*) from users where username = NEW.username and special_till > now()) = 1 THEN
    RETURN NEW;
  END IF;
  RETURN NULL;
END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER favorite_create BEFORE INSERT on favorite FOR EACH ROW EXECUTE PROCEDURE special_user_validation();

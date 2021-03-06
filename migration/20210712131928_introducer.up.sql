CREATE TABLE IF NOT EXISTS introducer
(
    username   varchar(255) primary key,
    introducer varchar(255),
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (introducer) REFERENCES users (username)
);

CREATE OR REPLACE FUNCTION increase_score() RETURNS trigger as
$$
  BEGIN
    UPDATE users SET score = score + 1 WHERE username = NEW.introducer;

    RETURN NEW;
  END;
$$
LANGUAGE 'plpgsql';

CREATE TRIGGER introducer_create BEFORE INSERT on introducer FOR EACH ROW EXECUTE PROCEDURE increase_score();

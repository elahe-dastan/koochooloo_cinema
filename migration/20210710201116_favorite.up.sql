CREATE TABLE IF NOT EXISTS favorite(
  user_id int,
  film_id int,
  name VARCHAR(255),
  primary key (user_id, film_id),
  FOREIGN KEY (user) REFERENCES registeration (id),
  FOREIGN KEY (film) REFERENCES film (id)
);


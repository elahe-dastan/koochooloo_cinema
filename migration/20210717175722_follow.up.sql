CREATE TABLE IF NOT EXISTS follow
(
    username  varchar(255),
    following varchar(255),
    primary key (username, following),
    FOREIGN KEY (username) REFERENCES users (username),
    FOREIGN KEY (following) REFERENCES users (username)
)
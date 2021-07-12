CREATE TABLE IF NOT EXISTS introducer
(
    username   varchar(255) primary key,
    introducer varchar(255),
    FOREIGN KEY (username) REFERENCES user (username),
    FOREIGN KEY (introducer) REFERENCES user (username)
)
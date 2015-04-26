--
-- sqlite3
--

CREATE TABLE user (
    user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
    email      VARCHAR(250) DEFAULT '',
    password   VARCHAR(250) DEFAULT NULL
);

CREATE EXTENSION IF NOT EXISTS citext;

CREATE DOMAIN Email AS citext
  CHECK ( value ~ '^[a-zA-Z0-9.!#$%&''*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$' );

DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users(
    nickname TEXT PRIMARY KEY CONSTRAINT nickname_right CHECK(nickname ~ '^[A-Za-z0-9.!#$%&''*+=?^_`{|}~-]{3,20}$'),
    email Email NOT NULL UNIQUE,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    password_digest TEXT NOT NULL,
    registration_time TIMESTAMP NOT NULL
);

-- in the end in config
SET timezone = 'Europe/Moscow';
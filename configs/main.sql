CREATE EXTENSION IF NOT EXISTS CITEXT;

DROP TABLE IF EXISTS Users CASCADE;
CREATE TABLE Users (
    nickname TEXT PRIMARY KEY CONSTRAINT nick_right CHECK(nickname ~ '^[A-Za-z0-9_\.]*$'),
    password_digest TEXT NOT NULL,
    email CITEXT NOT NULL UNIQUE CONSTRAINT email_right CHECK(email ~ '^.*@[A-Za-z0-9\-_\.]*$'),
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL
);
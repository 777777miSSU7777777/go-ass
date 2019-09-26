CREATE DATABASE goass;

USE goass;

CREATE TABLE authors (
    id BIGINT NOT NULL AUTO AUTO_INCREMENT,
    name NVARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY(id),
    CHECK(NAME <> "")
);

CREATE TABLE audio (
    id BIGINT NOT NULL AUTO_INCREMENT,
    author_id BIGINT NOT NULL,
    title NVARCHAR(50),
    PRIMARY KEY(id),
    FOREIGN KEY(author_id)
        REFERENCES authors(id)
);

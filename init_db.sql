CREATE DATABASE goass;

USE goass;

CREATE TABLE authors (
    id BIGINT NOT NULL AUTO AUTO_INCREMENT,
    name NVARCHAR(50) NOT NULL UNIQUE,
    PRIMARY KEY(id),
    CHECK(name <> "")
);

CREATE TABLE audio (
    id BIGINT NOT NULL AUTO_INCREMENT,
    author NVARCHAR(50) NOT_NULL,
    title NVARCHAR(50) NOT NULL,
    PRIMARY KEY(id),
    CHECK(author <> ""),
    CHECK(title <> "")
);

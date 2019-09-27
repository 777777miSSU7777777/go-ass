CREATE DATABASE goass;

USE goass;

CREATE TABLE audio (
    id BIGINT NOT NULL AUTO_INCREMENT,
    author NVARCHAR(50) NOT NULL,
    title NVARCHAR(50) NOT NULL,
    PRIMARY KEY(id),
    CHECK(author <> ""),
    CHECK(title <> "")
);

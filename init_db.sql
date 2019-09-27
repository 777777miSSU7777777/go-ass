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

CREATE TABLE tables_last_id (
    table_name NVARCHAR(50) NOT NULL,
    last_id BIGINT NOT NULL,
    PRIMARY KEY(table_name)
);

INSERT INTO tables_last_id(table_name, last_id) VALUES ("audio", 0);

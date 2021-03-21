-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
       id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
       name CHAR(100) UNIQUE,
       email CHAR(255) UNIQUE,
       creation_date TIMESTAMP,
       CONSTRAINT contacts_unique UNIQUE (name, email));
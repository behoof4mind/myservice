-- +migrate Up
CREATE database IF NOT EXISTS myservice;

CREATE TABLE IF NOT EXISTS myservice.users (
   id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
   name CHAR(100) UNIQUE,
    email CHAR(255) UNIQUE,
    creation_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    CONSTRAINT contacts_unique UNIQUE (name, email));
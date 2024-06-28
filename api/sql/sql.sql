CREATE DATABASE IF NOT EXISTS devbook;

USE devbook

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id INTEGER AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    nick VARCHAR(255) NOT NULL unique,
    email VARCHAR(255) NOT NULL unique,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
) ENGINE=INNODB;
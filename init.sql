CREATE DATABASE IF NOT EXISTS api_db;
USE api_db;

CREATE TABLE IF NOT EXISTS Persons (
     Name varchar(255),
     Age int,
     Balance double,
     Email varchar(255) NOT NULL,
     Address varchar(255),
     PRIMARY KEY (Email)
);

INSERT INTO Persons VALUES ('Bogdan Guranda', 25, 2500.50, "bogdan.g@gmail.com", "Cernauti 24, Oradea");
INSERT INTO Persons VALUES ('Andrei Guranda', 28, 34580, "andrei.g@gmail.com", "Mehedinti 4, Cluj");
INSERT INTO Persons VALUES ('Ricky Martin', 46, 34580, "ricky.m@gmail.com", "10th Street, London");
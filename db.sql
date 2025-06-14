-- Start with creating a database
-- Delete the database if it exists with the following command:
-- DROP DATABASE IF EXISTS test_db;
CREATE DATABASE IF NOT EXISTS test_db;
-- Actually use the database to run other commands
USE test_db;
-- Create a table with a primary key and a foreign key
CREATE TABLE IF NOT EXISTS test_table (
    id INT PRIMARY KEY UNIQUE NOT NULL DEFAULT (UUID()),
    FOREIGN KEY (id) REFERENCES test_table(id)
    name VARCHAR(255) NOT NULL,
    guild VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    
) ENGINE=InnoDB;


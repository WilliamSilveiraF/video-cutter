-- create_table.sql

CREATE TABLE IF NOT EXISTS persons (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    gender VARCHAR(50),
    contact VARCHAR(255),
    birthday DATE
);

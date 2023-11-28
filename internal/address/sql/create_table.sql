-- create_address_table.sql

CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    zip VARCHAR(20),
    street VARCHAR(255),
    unit VARCHAR(50),
    city VARCHAR(100),
    state VARCHAR(100)
);
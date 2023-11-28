-- create_table.sql

CREATE TABLE IF NOT EXISTS audio (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    filename VARCHAR(255) NOT NULL
);

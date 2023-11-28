-- create_card_table.sql

CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    reference VARCHAR(255) NOT NULL,
    validity VARCHAR(10) NOT NULL,
    cvv VARCHAR(5) NOT NULL,
    cardholder VARCHAR(255) NOT NULL
);
-- create_tables.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    use_terms_id INT,
    FOREIGN KEY (use_terms_id) REFERENCES use_terms(id)
);

CREATE TABLE IF NOT EXISTS use_terms (
    id SERIAL PRIMARY KEY,
    version VARCHAR(255) NOT NULL,
    description TEXT NOT NULL
);
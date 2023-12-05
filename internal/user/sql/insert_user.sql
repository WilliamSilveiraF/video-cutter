-- insert_user.sql

INSERT INTO users (email, password, use_terms_id) VALUES ($1, $2, $3) RETURNING id;

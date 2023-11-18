-- insert_user.sql

INSERT INTO users (email, password) VALUES ($1, $2);

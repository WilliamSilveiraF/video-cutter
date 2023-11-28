-- insert_address.sql

INSERT INTO addresses (user_id, zip, street, unit, city, state)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
-- update_address.sql

UPDATE addresses
SET zip = $2, street = $3, unit = $4, city = $5, state = $6
WHERE user_id = $1;

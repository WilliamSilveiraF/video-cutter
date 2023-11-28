-- retrieve_address.sql

SELECT id, user_id, zip, street, unit, city, state
FROM addresses
WHERE user_id = $1;

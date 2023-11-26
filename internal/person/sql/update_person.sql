-- update_person.sql

UPDATE persons
SET user_id = $2, first_name = $3, last_name = $4, gender = $5, contact = $6, birthday = $7
WHERE id = $1;

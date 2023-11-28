-- update_person.sql

UPDATE persons
SET first_name = $2, last_name = $3, gender = $4, contact = $5, birthday = $6
WHERE user_id = $1;

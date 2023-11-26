-- insert_person.sql

INSERT INTO persons (user_id, first_name, last_name, gender, contact, birthday)
VALUES ($1, $2, $3, $4, $5, $6);
-- retrieve_person.sql

SELECT id, user_id, first_name, last_name, gender, contact, birthday
FROM persons
WHERE user_id = $1;

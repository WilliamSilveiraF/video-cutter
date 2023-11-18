-- retrieve_user.sql

SELECT password FROM users WHERE email = $1;

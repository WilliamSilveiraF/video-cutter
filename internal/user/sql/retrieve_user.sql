-- retrieve_user.sql

SELECT id, email, password FROM users WHERE email = $1;

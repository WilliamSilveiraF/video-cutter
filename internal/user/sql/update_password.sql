-- update_password.sql

UPDATE users SET password = $1 WHERE email = $2;

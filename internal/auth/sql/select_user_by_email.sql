-- select_user_by_email.sql

SELECT id, password FROM users WHERE email = $1;

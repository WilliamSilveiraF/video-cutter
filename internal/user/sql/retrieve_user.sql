-- retrieve_user.sql

SELECT id, email, password, use_terms_id FROM users WHERE email = $1;

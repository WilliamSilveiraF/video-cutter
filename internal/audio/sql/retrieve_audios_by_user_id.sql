-- retrieve_audios_by_user_id.sql

SELECT id, user_id, filename FROM audio WHERE user_id = $1;

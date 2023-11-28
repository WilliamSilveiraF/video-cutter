-- retrieve_audios_by_user_id.sql

SELECT id, filename, transcription FROM audio WHERE user_id = $1;

-- insert_audio.sql

INSERT INTO audio (user_id, filename, transcription) VALUES ($1, $2, $3) RETURNING id;

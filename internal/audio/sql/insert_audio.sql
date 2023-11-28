-- insert_audio.sql

INSERT INTO audio (user_id, filename) VALUES ($1, $2) RETURNING id;

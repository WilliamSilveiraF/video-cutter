-- delete_card_by_id.sql

DELETE FROM cards
WHERE id = $1;

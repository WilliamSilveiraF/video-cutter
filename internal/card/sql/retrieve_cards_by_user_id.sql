-- retrieve_cards_by_user_id.sql

SELECT id, reference, validity, cvv, cardholder
FROM cards
WHERE user_id = $1;

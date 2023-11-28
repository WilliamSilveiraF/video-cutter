-- insert_card.sql

INSERT INTO cards (user_id, reference, validity, cvv, cardholder)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;
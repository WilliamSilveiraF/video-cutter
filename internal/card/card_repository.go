package card

import (
    "workflow-editor/db"
)

func InsertCard(card Card) (int, error) {
    sqlQuery, err := db.ReadSQLFile("internal/card/sql/insert_card.sql")
    if err != nil {
        return 0, err
    }

    var cardID int
    err = db.GetDB().QueryRow(sqlQuery, card.UserID, card.Reference, card.Validity, card.CVV, card.Cardholder).Scan(&cardID)
    if err != nil {
        return 0, err
    }

    return cardID, nil
}

func RetrieveCardsByUserID(userID int) ([]Card, error) {
    sqlQuery, err := db.ReadSQLFile("internal/card/sql/retrieve_cards_by_user_id.sql")
    if err != nil {
        return nil, err
    }

    rows, err := db.GetDB().Query(sqlQuery, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cards []Card
    for rows.Next() {
        var card Card
        if err := rows.Scan(&card.ID, &card.Reference, &card.Validity, &card.CVV, &card.Cardholder); err != nil {
            return nil, err
        }
        cards = append(cards, card)
    }

    return cards, nil
}

func DeleteCardByID(cardID int) error {
    sqlQuery, err := db.ReadSQLFile("internal/card/sql/delete_card_by_id.sql")
    if err != nil {
        return err
    }

    _, err = db.GetDB().Exec(sqlQuery, cardID)
    return err
}

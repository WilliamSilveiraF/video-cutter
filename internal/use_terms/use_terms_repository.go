package use_terms

import (
    "workflow-editor/db"
)

func GetUseTermsByID(id int) (*UseTerms, error) {
    var useTerms UseTerms
    sqlQuery, err := db.ReadSQLFile("internal/use_terms/sql/retrieve_use_terms.sql")
    if err != nil {
        return nil, err
    }

    row := db.GetDB().QueryRow(sqlQuery, id)
    err = row.Scan(&useTerms.ID, &useTerms.Version, &useTerms.Description)
    if err != nil {
        return nil, err
    }

    return &useTerms, nil
}

func GetLatestUseTermsID() (int, error) {
    var id int
    sqlQuery := "SELECT id FROM use_terms ORDER BY id DESC LIMIT 1"
    err := db.GetDB().QueryRow(sqlQuery).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}
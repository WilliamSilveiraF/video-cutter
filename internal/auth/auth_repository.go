package auth

import (
	"workflow-editor/db"
)


func InsertUser(user User) error {
	sqlQuery, err := db.ReadSQLFile("internal/auth/sql/insert_user.sql")
	if err != nil {
		return err
	}

	stmt, err := db.GetDB().Prepare(sqlQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
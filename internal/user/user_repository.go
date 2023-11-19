package user

import (
	"workflow-editor/db"
)

func RetrieveUser(email string) (*User, error) {
	var user User

	sqlQuery, err := db.ReadSQLFile("internal/user/sql/retrieve_user.sql")
	if err != nil {
		return nil, err
	}

	row := db.GetDB().QueryRow(sqlQuery, email)
	err = row.Scan(&user.Password)
	if err != nil {
		return nil , err
	}

	user.Email = email
	return &user, nil
}

func InsertUser(user User) error {
	sqlQuery, err := db.ReadSQLFile("internal/user/sql/insert_user.sql")
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
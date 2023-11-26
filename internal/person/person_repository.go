package person

import (
	"workflow-editor/db"
)

func RetrievePerson(id int) (*Person, error) {
	var person Person
	
	sqlQuery, err := db.ReadSQLFile("internal/person/sql/retrieve_person.sql")
	row := db.GetDB().QueryRow(sqlQuery, id)
	err = row.Scan(&person.ID, &person.UserID, &person.FirstName, &person.LastName, &person.Gender, &person.Contact, &person.Birthday)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

func InsertPerson(person Person) error {
	sqlQuery, err := db.ReadSQLFile("internal/person/sql/insert_person.sql")
	
	if err != nil {
		return err
	}

	stmt, err := db.GetDB().Prepare(sqlQuery)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	_, err = stmt.Exec(person.UserID, person.FirstName, person.LastName, person.Gender, person.Contact, person.Birthday)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePerson(person Person) error {
	sqlQuery, err := db.ReadSQLFile("internal/person/sql/update_person.sql")
	if err != nil {
		return err
	}

	stmt, err := db.GetDB().Prepare(sqlQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(person.ID, person.UserID, person.FirstName, person.LastName, person.Gender, person.Contact, person.Birthday)
	if err != nil {
		return err
	}

	return nil
}

package repository

import (
	"database/sql"
	"rest_api_deploy/structs"
)

func GetAllPerson(db *sql.DB) (results []structs.Person, err error) {
	query := "SELECT * FROM person"

	rows, err := db.Query(query)
	if err != nil {
    panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var person = structs.Person{}

		err = rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		
		if err != nil {
			panic(err)
		}

		results = append(results, person)
	}

	return results, nil
}

func InsertPerson(db *sql.DB, person structs.Person) (err error) {
	query :=
		`INSERT INTO person (id, first_name, last_name)
		VALUES ($1, $2, $3)
		Returning *`

	errs := db.QueryRow(query, person.ID, person.FirstName, person.LastName)

	return errs.Err()
}

func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
	query :=
		`UPDATE person 
		SET first_name = $1, last_name = $2
		WHERE id = $3`

	errs := db.QueryRow(query, person.ID, person.FirstName, person.LastName)

	return errs.Err()
}

func deletePerson(db *sql.DB, person structs.Person) (err error) {
	query :=
		`DELETE FROM person 
		WHERE id = $1`

	errs := db.QueryRow(query, person.ID)

	return errs.Err()
}
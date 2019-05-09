package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	db := OpenDB()
	return db
}

func OpenDB() *sql.DB {
	database, err := sql.Open("sqlite3", "./skynet.db")
	if err != nil {
		log.Fatalf("Panic. Open DB Fail: %s", err)
		return nil
	}

	return database
}

func CreateDB(db *sql.DB) error {
	table1 := "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)"
	table2 := "CREATE TABLE IF NOT EXISTS product (id INTEGER PRIMARY KEY, name TEXT, price TEXT)"

	statement1, err := db.Prepare(table1)
	if err != nil {
		return err
	}

	statement2, err := db.Prepare(table2)
	if err != nil {
		return err
	}

	_, err = statement1.Exec()
	if err != nil {
		return err
	}

	_, err = statement2.Exec()
	if err != nil {
		return err
	}

	return nil
}

func Update() {

}

func QueryByName(db *sql.DB, name string) (string, error) {

	rows, err := db.Query("SELECT id, firstname, lastname FROM people WHERE firstname = '" + name + "'")

	if err != nil {
		return "", err
	}

	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
	}

	return firstname, nil
}

func Query(db *sql.DB) {

	rows, _ := db.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		// fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)

	}
}

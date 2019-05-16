package sql

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func GetDB() *sql.DB {
	db := openDB("./skynet.db")
	return db
}

func openDB(dbName string) *sql.DB {
	database, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Panic. Open DB Fail: %s", err)
		return nil
	}

	return database
}

// CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)
func CreateDB(db *sql.DB, objects []interface{}) error {

	var createTableQuery []string

	for _, object := range objects {

		createQuery := ""

		if reflect.ValueOf(object).Kind() == reflect.Struct {
			tableName := reflect.TypeOf(object).Name()
			createQuery = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT", tableName)

			v := reflect.ValueOf(object)

			for i := 0; i < v.NumField(); i++ {

				fieldName := v.Type().Field(i).Name
				createQuery = fmt.Sprintf("%s, %s TEXT", createQuery, fieldName)

			}

			createQuery = fmt.Sprintf("%s)", createQuery)
		}

		createTableQuery = append(createTableQuery, createQuery)
	}

	for _, table := range createTableQuery {

		fmt.Println(table)

		statement, err := db.Prepare(table)
		if err != nil {
			return err
		}

		_, err = statement.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}

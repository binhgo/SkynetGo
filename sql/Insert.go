package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
)

func Insert(db *sql.DB, object interface{}) error {
	defer db.Close()

	insertQuery, err := CreateInsertQuery(object)
	if err != nil {
		return err
	}

	statement, _ := db.Prepare(insertQuery)
	statement.Exec()
	return nil
}

func CreateInsertQuery(q interface{}) (string, error) {

	if reflect.ValueOf(q).Kind() == reflect.Struct {

		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("insert into %s values(", t)

		v := reflect.ValueOf(q)

		for i := 0; i < v.NumField(); i++ {

			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return "", errors.New("Unsupported type in struct")
			}
		}

		query = fmt.Sprintf("%s)", query)
		return query, nil
	}

	return "", errors.New("Unsupported type in struct")
}

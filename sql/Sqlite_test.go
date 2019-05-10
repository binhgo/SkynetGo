package sql

import (
	"testing"
)

type Message struct {
	id      string
	content string
	email   string
}

type Product struct {
	id    string
	name  string
	price string
}

func TestOpenDB(t *testing.T) {
	db := GetDB()
	defer db.Close()

	if db == nil {
		t.Errorf("Error TestOpenDB")
	}
}

func TestCreateDB(t *testing.T) {
	db := GetDB()
	defer db.Close()

	var product Product
	var mess Message

	var tables []interface{}

	tables = append(tables, product, mess)

	err := CreateDB(db, tables)

	if err != nil {
		t.Errorf("Error TestCreateDB")
	}
}

// func TestQueryString_From(t *testing.T) {
// 	//assert := assert.New(t)
// 	builder := &Builder{}
// 	get := builder.Select([]string{"*"}).
// 		From("people").
// 		Where("firstname", EQUAL, "Ric").
// 		AndWhere("lastname", EQUAL, "abc").
// 		OrderBy("lastname", DESC).
// 		Get()
//
// 	fmt.Println(get)
//
// 	//assert.Equal("SELECT * FROM people", get)
// }

// func TestInsert(t *testing.T) {
// 	assert := assert.New(t)
//
// 	db := OpenDB()
// 	err := CreateDB(db)
//
// 	if err != nil {
// 		t.Errorf("Error TestCreateDB")
// 	}
//
// 	Insert(db)
//
// 	name, err := QueryByName(db, "Nic")
// 	assert.Nil(err)
// 	assert.Equal("Nic", name)
// }

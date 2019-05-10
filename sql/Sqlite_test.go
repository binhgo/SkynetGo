package sql

import (
	"fmt"
	"testing"
)

func TestOpenDB(t *testing.T) {
	db := OpenDB()

	if db == nil {
		t.Errorf("Error TestOpenDB")
	}
}

func TestCreateDB(t *testing.T) {
	db := OpenDB()
	err := CreateDB(db)

	if err != nil {
		t.Errorf("Error TestCreateDB")
	}
}

func TestQueryString_From(t *testing.T) {
	//assert := assert.New(t)

	builder := &QueryBuilder{}
	get := builder.Select().
		From("people").
		Where("firstname", EQUAL, "Nic").
		OrderBy("lastname", DESC)

	fmt.Println(get)

	db := OpenDB()
	//Insert(db)
	get.Exec(db)

	//assert.Equal("select * from people", get)
}

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

package sql

import (
	"database/sql"
	"fmt"
)

// SELECT * FROM table WHERE key = 'value' and key = 'value' order by

type OrderType string

const (
	ASC  OrderType = "ASC"
	DESC OrderType = "DESC"
)

type Condition string

const (
	EQUAL   Condition = "="
	GREATER Condition = ">"
	LESS    Condition = "<"
	DIFF    Condition = "<>"
)

type QueryBuilder struct {
	Query string
}

func (builder *QueryBuilder) Select() *QueryBuilder {
	builder.Query = fmt.Sprintf("select *")
	return builder
}

func (builder *QueryBuilder) From(table string) *QueryBuilder {
	builder.Query = fmt.Sprintf("%s from %s", builder.Query, table)
	return builder
}

func (builder *QueryBuilder) Where(column string, condition Condition, value string) *QueryBuilder {
	builder.Query = fmt.Sprintf("%s where %s %s '%s'", builder.Query, column, condition, value)
	return builder
}

func (builder *QueryBuilder) AndWhere(column string, condition Condition, value string) *QueryBuilder {
	builder.Query = fmt.Sprintf("%s and %s %s '%s'", builder.Query, column, condition, value)
	return builder
}

func (builder *QueryBuilder) OrWhere(column string, condition Condition, value string) *QueryBuilder {
	builder.Query = fmt.Sprintf("%s or %s = '%s'", builder.Query, column, value)
	return builder
}

func (builder *QueryBuilder) OrderBy(column, orderType OrderType) *QueryBuilder {
	builder.Query = fmt.Sprintf("%s order by %s %s", builder.Query, column, orderType)
	return builder
}

func (builder *QueryBuilder) Exec(db *sql.DB) (*sql.Rows, error) {

	rows, err := db.Query(builder.Query)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

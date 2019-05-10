package sql

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"strings"
)

// SELECT * FROM table WHERE key = 'value' and key = 'value' order by

const (
	LIKE    Condition = "LIKE"
	BETWEEN Condition = "BETWEEN"
)

type Operator string

const (
	AND Operator = "AND"
	OR  Operator = "OR"
)

type Builder struct {
	SelectQuery    []string
	FromQuery      string
	WhereQuery     map[Operator][]string
	OrderByQuery   string
	OrderTypeQuery OrderType
	GroupByQuery   []string
	HavingQuery    []string
	LimitQuery     int
	OffsetQuery    int
}

func (builder *Builder) Select(selects []string) *Builder {
	builder.SelectQuery = selects
	return builder
}

func (builder *Builder) From(table string) *Builder {
	builder.FromQuery = table
	return builder
}

func (builder *Builder) Where(column string, condition Condition, value string) *Builder {
	if builder.WhereQuery == nil {
		builder.WhereQuery = make(map[Operator][]string)
	}
	str := fmt.Sprintf("`%s` %s '%s'", column, condition, value)
	//builder.WhereQuery = append(builder.WhereQuery, map[Operator][]string{AND: []string{str}})
	if _, ok := builder.WhereQuery[AND]; !ok {
		builder.WhereQuery[AND] = []string{}
	}
	builder.WhereQuery[AND] = append(builder.WhereQuery[AND], str)
	return builder
}

func (builder *Builder) AndWhere(column string, condition Condition, value string) *Builder {
	return builder.Where(column, condition, value)
}

func (builder *Builder) OrWhere(column string, condition Condition, value string) *Builder {
	if builder.WhereQuery == nil {
		builder.WhereQuery = make(map[Operator][]string)
	}
	str := fmt.Sprintf("`%s` %s '%s'", column, condition, value)
	if _, ok := builder.WhereQuery[OR]; !ok {
		builder.WhereQuery[OR] = []string{}
	}
	builder.WhereQuery[OR] = append(builder.WhereQuery[OR], str)
	return builder
}

func (builder *Builder) OrderBy(column string, orderType OrderType) *Builder {
	builder.OrderByQuery = column
	builder.OrderTypeQuery = orderType
	return builder
}

func (builder *Builder) GroupBy(groups []string) *Builder {
	builder.GroupByQuery = groups
	return builder
}

func (builder *Builder) Having(having []string) *Builder {
	builder.HavingQuery = having
	return builder
}

func (builder *Builder) Limit(limit int) *Builder {
	builder.LimitQuery = limit
	return builder
}

func (builder *Builder) Offset(offset int) *Builder {
	builder.OffsetQuery = offset
	return builder
}

func (builder *Builder) Get() string {
	var sql string
	if len(builder.SelectQuery) > 0 {
		sql = fmt.Sprintf("SELECT %s", strings.Join(builder.SelectQuery, ","))
	}else{
		sql = "SELECT *"
	}
	if builder.FromQuery == "" {
		log.Fatal("From is required")
	}
	sql = fmt.Sprintf("%s FROM %s", sql, builder.FromQuery)
	if len(builder.WhereQuery) > 0 {
		where := []string{}
		if len(builder.WhereQuery[AND]) > 0 {
			where = append(where, strings.Join(builder.WhereQuery[AND], " AND "))
		}
		if len(builder.WhereQuery[OR]) > 0 {
			where = append(where, strings.Join(builder.WhereQuery[OR], " OR "))
		}
		sql = fmt.Sprintf("%s WHERE %s", sql, strings.Join(where, " AND "))
	}

	return sql
}

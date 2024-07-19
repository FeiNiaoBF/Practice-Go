package clause

import (
	"fmt"
	"foorm/log"
	"os"
	"strings"
)

// generator 实现各个子句的生成规则

type generator func(values ...any) (string, []any)

var generators map[Type]generator

func init() {
	generators = make(map[Type]generator)
	// init type
	generators[INSERT] = _insert
	generators[VALUES] = _values
	generators[SELECT] = _select
	generators[LIMIT] = _limit
	generators[WHERE] = _where
	generators[ORDERBY] = _orderBy
	generators[UPDATE] = _update
	generators[DELETE] = _delete
	generators[COUNT] = _count
}

// genBindVars return a string of num value ? of ASCII
// num is 2, s is "?, ?"
func genBindVars(num int) string {
	var vars = make([]string, 0, num)
	for i := 0; i < num; i++ {
		vars = append(vars, "?")
	}
	return strings.Join(vars, ", ")
}

// _insert is build insert clause
func _insert(values ...any) (string, []any) {
	// INSERT INTO $TABLE (fields) ...
	lognil(values)
	table := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("INSERT INTO %s (%v)", table, fields), []any{}
}

// _value
func _values(values ...any) (string, []any) {
	lognil(values)
	// VALUES ($v1), ($v2)...
	// "('Jack', 18), ('Tom', 20)"
	var vars []any
	var builder strings.Builder
	var bindStr string

	builder.WriteString("VALUES ")
	for i, val := range values {
		v := val.([]any)

		if bindStr == "" {
			bindStr = genBindVars(len(v))

		}
		builder.WriteString(fmt.Sprintf("(%v)", bindStr))

		if i+1 != len(values) {
			builder.WriteString(", ")
		}
		vars = append(vars, v...)
	}

	return builder.String(), vars
}

func _select(values ...any) (string, []any) {
	// SELECT $fields FROM $tableName
	lognil(values)
	tableName := values[0]
	fields := strings.Join(values[1].([]string), ", ")
	return fmt.Sprintf("SELECT %v FROM %s", fields, tableName), []any{}
}

func _limit(values ...any) (string, []any) {
	// LIMIT $num
	lognil(values)
	return "LIMIT ?", values
}

func _where(values ...any) (string, []any) {
	// WHERE $desc
	lognil(values)
	desc, vars := values[0], values[1:]
	return fmt.Sprintf("WHERE %s", desc), vars
}

func _orderBy(values ...any) (string, []any) {
	return fmt.Sprintf("ORDER BY %s", values[0]), []any{}
}

func _update(values ...any) (string, []any) {
	tableName := values[0]
	// update 只允许变量为map[string]any
	// map: {("Name": "Jack"), ("Age": 20)}
	m := values[1].(map[string]any)
	var keys []string
	var vars []any
	for k, v := range m {
		keys = append(keys, k+"=?")
		vars = append(vars, v)
	}
	return fmt.Sprintf("UPDATE %s SET %s", tableName, strings.Join(keys, ", ")), vars
}

func _delete(values ...any) (string, []any) {
	return fmt.Sprintf("DELETE FROM %s", values[0]), []any{}
}

func _count(values ...any) (string, []any) {
	return _select(values[0], []string{"count(*)"})
}

func lognil(values ...any) {
	if len(values) == 0 {
		log.Error("values is empty")
		os.Exit(2)
	}
}

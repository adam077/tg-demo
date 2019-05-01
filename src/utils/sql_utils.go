package utils

import (
	"fmt"
	"github.com/json-iterator/go"
	"strconv"
	"strings"
)

func GetStr(value *string) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("'%s'", strings.Replace(*value, "'", "''", -1))
}

func GetInt(value *int) string {
	if value == nil {
		return "null"
	}
	return strconv.Itoa(*value)
}

func GetFloat(value *float64) string {
	if value == nil {
		return "null"
	}
	return strconv.FormatFloat(*value, 'f', 2, 32)
}

func GetRawMessage(value *jsoniter.RawMessage) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("'%s'", strings.Replace(string(*value), "'", "''", -1))
}

func GetBool(value *bool) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprintf("'%s'", strings.Replace(strconv.FormatBool(*value), "'", "''", -1))
}

func CreateBatchSql(tb string, cols []string, values [][]string, keys, updateCols []string) string {
	valuesBatch := make([]string, 0)
	for x := range values {
		valuesBatch = append(valuesBatch, fmt.Sprintf("(%s)", strings.Join(values[x], ",")))
	}
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s on conflict (%s) do update set %s;",
		tb, strings.Join(cols, ","), strings.Join(valuesBatch, ","), strings.Join(keys, ","), strings.Join(updateCols, ","))
}

func CreateSql(tb string, cols, values, keys, updateCols []string) string {
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) on conflict (%s) do update set %s;",
		tb, strings.Join(cols, ","), strings.Join(values, ","), strings.Join(keys, ","), strings.Join(updateCols, ","))
}

func UpdateForm(col string) string {
	return col + " = excluded." + col
}

func GetUpdateTail(cols []string) []string {
	a := make([]string, 0)
	for x := range cols {
		if cols[x] != "create_time" {
			a = append(a, UpdateForm(cols[x]))
		}
	}
	return a
}

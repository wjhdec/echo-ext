package db

import "strings"

type WhereMap map[string]any

func (w WhereMap) AndWhere() (whereStr string, values []any) {
	keys := make([]string, 0, len(w))
	values = make([]any, 0, len(w))
	for key, value := range w {
		keys = append(keys, key)
		values = append(values, value)
	}
	whereStr = strings.Join(keys, " and ")
	return
}

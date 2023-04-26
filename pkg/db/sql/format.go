package sql

import (
	"regexp"
)

// Clear 刷新sql
func Clear(sql string) string {
	if sql == "" {
		return ""
	}
	r, _ := regexp.Compile(`[\r\n\t]`)
	sql2 := r.ReplaceAllString(sql, " ")
	r2, _ := regexp.Compile(`\s{2,}`)
	return r2.ReplaceAllString(sql2, " ")
}

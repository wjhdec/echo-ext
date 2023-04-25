package custime

import (
	"fmt"
	"strings"
	"time"
)

const defaultFmt = "2006-01-02 15:04:05"

type FormatTime time.Time

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(defaultFmt, timeStr)
	*t = FormatTime(t1)
	return err
}

func (t *FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(*t).Format(defaultFmt))
	return []byte(formatted), nil
}

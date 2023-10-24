package custime

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const defaultFmt = "2006-01-02 15:04:05"

type FormatTime struct {
	time.Time
}

func Now() FormatTime {
	return FormatTime{time.Now()}
}

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	if string(data) == "" || string(data) == "null" {
		return nil
	}
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse(defaultFmt, timeStr)
	t.Time = t1
	return errors.WithStack(err)
}

func (t *FormatTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(defaultFmt))
	return []byte(formatted), nil
}

func (t *FormatTime) Scan(value any) error {
	switch tv := value.(type) {
	case time.Time:
		t.Time = tv
	case FormatTime:
		*t = tv
	default:
		return errors.Errorf("cannot scan type%T into FormatTime", value)
	}
	return nil
}

func (t FormatTime) Value() (driver.Value, error) {
	return t.Time, nil
}

package dbtype

import (
	"encoding/json"
	"fmt"
)

type JSON struct {
	json.RawMessage
}

func (j *JSON) Scan(value any) (err error) {
	if value == nil {
		return
	}
	switch v := value.(type) {
	case []byte:
		j.RawMessage = v
	case string:
		j.RawMessage = []byte(v)
	default:
		err = fmt.Errorf("cannot sql.Scan() JSONField from: %#v", value)
	}
	return
}

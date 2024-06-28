package location

import (
	"database/sql"
	"fmt"
	"strings"
)

type Location struct {
	Id          string         `db:"id"`
	Name        string         `db:"name"`
	ParkingLots StringArray    `db:"parking_lots"`
	Address     sql.NullString `db:"address"`
}

type StringArray []string

// I need this because postgress array is like {value1,value2}
func (a *StringArray) Scan(src interface{}) error {
	var srcStr string
	switch src := src.(type) {
	case []byte:
		srcStr = string(src)
	case string:
		srcStr = src
	default:
		return fmt.Errorf("unsupported scan type %T", src)
	}

	srcStr = strings.Trim(srcStr, "{}")
	if srcStr == "" {
		*a = []string{}
		return nil
	}

	*a = strings.Split(srcStr, ",")
	return nil
}

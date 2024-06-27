package location

import "database/sql"

type Location struct {
	Id           string
	Name         string
	Parking_lots []string
	Address      sql.NullString
}

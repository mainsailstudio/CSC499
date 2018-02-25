package dbinfo

import (
	dynauthconst "dynauthconst"
)

// Db returns a slice with the needed db information
func Db() []string {
	dbinfo := []string{"mysql", dynauthconst.DatabaseUser + ":" + dynauthconst.DatabasePass + "@/" + dynauthconst.DatabaseName}
	return dbinfo
}

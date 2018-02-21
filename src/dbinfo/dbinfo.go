package dbinfo

// Db returns a slice with the needed db information
func Db() []string {
	dbinfo := []string{"mysql", "root:root@/dynauth_dev"}
	return dbinfo
}

package database

var connetion string

func init() {
	connetion = "PostgreSQL"
}

func GetDB() string {
	return connetion
}

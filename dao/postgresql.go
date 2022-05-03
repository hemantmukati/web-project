package dao

import (
	"database/sql"
	"fmt"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
)

// DB set up
func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}

	return db
}

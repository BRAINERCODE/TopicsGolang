package Database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetMySqlClient() (db *sql.DB, e error) {

	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/products")

	if err != nil {
		return nil, err
	}

	return db, nil
}

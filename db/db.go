package db

import (
	"database/sql"
	"fmt"
	"os"

	sqlbuilder "github.com/huandu/go-sqlbuilder" // SQL builder
	_ "github.com/lib/pq"
)

// DBConnection is a holder for a database connection
type DBConnection struct {
	dbType  string
	DB      *sql.DB
	Builder sqlbuilder.Flavor
	IDMax   int
}

// NewDBConnection connects to the DB
func NewDBConnection() (conn *DBConnection, err error) {
	var db_host string
	conn = new(DBConnection)
	conn.dbType = "postgres"
	if os.Getenv("DB_HOST") == "" {
		db_host = "postgres"
	} else {
		db_host = os.Getenv("DB_HOST")
	}

	conn.DB, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@%s:5432/postgres?sslmode=disable", db_host))
	if err != nil {
		return
	}
	err = conn.DB.Ping()
	if err != nil {
		return
	}
	conn.IDMax = 9223372036854775807

	conn.Builder = sqlbuilder.PostgreSQL
	return
}

func (conn *DBConnection) Close() {
	conn.DB.Close()
}

package dbconnector

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

var conn *sql.DB

func init() {
	connect()
}

func connect() (*sql.DB, error) {
	var err error
	conn, err = sql.Open("mssql", "mssql://localhost:1433; user id=sa;password=tempPass123")
	if err != nil {
		return nil, &connectionDriverError{}
	}
	err = conn.Ping()
	if err != nil {
		return nil, &connectionError{}
	}
	return conn, nil
}

func Get() (*sql.DB, error) {
	if conn != nil {
		return conn, nil
	}
	return connect()
}

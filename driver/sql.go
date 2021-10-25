package driver

import (
	"database/sql"
	"fmt"

	"example.com/sample-api/errors"

	_ "github.com/go-sql-driver/mysql"
)

type Configs struct {
	Host     string
	Username string
	Password string
	Port     int
	Database string
}

func InitializeDB(c Configs) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Username, c.Password, c.Host, c.Port, c.Database)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

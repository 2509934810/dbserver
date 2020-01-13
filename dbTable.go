package dbserver

import (
	"database/sql"

	"github.com/labstack/echo"
)

type Db struct {
	db *sql.DB
}

func (db Db) CreateTable(e echo.Context) error {
	
}

package dbserver

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"net/http"
)

func CreateDb(c echo.Context) error {
	dbName := c.QueryParam("dbname")
	sqlQuery := "create database " + dbName + ";"
	db, err := sql.Open("mysql", "root:lei123@/lei")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	return c.String(http.StatusOK, "you have create a database")
}

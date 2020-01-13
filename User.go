package dbserver

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func CreateUser(c echo.Context) error {
	name := c.QueryParam("name")
	password := c.QueryParam("password")
	db, err := sql.Open("mysql", "root:lei123@/lei")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	stmt, err := db.Prepare("insert into squarenum values(?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, password)
	if err != nil {
		panic(err.Error())
	}
	return c.HTML(http.StatusCreated, "<b>"+name+"data have been created</b>")
}

// func showUser(c echo.Context) error {
// 	id := c.QueryParam("Id")
// 	db, err := sql.Open("mysql", "root:lei123@/lei")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	rows, err := db.Query("select * from squarenum where Id = %d", id)
// 	value := make([]sql.Rows, len(rows.Columns()))

// }

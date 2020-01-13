package main

import (
	"io"
	"net/http"
	"os"

	"github.com/2509934810/dbserver"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func login(c echo.Context) error {
	name := c.QueryParam("name")
	password := c.QueryParam("password")
	return c.String(http.StatusOK, name+password)
}

func loginForm(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")
	return c.String(http.StatusOK, name+password)
}

type Person struct {
	name string
	age  int8
}

func show(c echo.Context) error {
	a := new(Person)
	if err := c.Bind(a); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, a)
}

func save(c echo.Context) error {
	name := c.FormValue("name")
	avator, err := c.FormFile("avator")
	if err != nil {
		panic(err.Error())
	}
	src, err := avator.Open()
	if err != nil {
		panic(err.Error())
	}
	defer src.Close()
	dst, err := os.Create("/Users/jlei-ext/go/" + avator.Filename)
	if err != nil {
		panic(err.Error())
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return c.HTML(http.StatusOK, "<b>Thanks"+name+"contribute your file</b>")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.GET("/login", login)
	e.GET("/user/:id", getUser)
	e.POST("/loginform", loginForm)
	e.POST("/save", save)
	e.POST("/show", show)
	e.GET("/createUser", dbserver.CreateUser)
	e.GET("/database/create", dbserver.CreateDb)

	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.Start(":1234")
}

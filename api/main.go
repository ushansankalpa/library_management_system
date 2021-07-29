package main

import (
	"net/http"
	"new_library/controler"
	"new_library/env"
	"new_library/model"

	"new_library/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	env.LoadEnvs()

	database0, err := gorm.Open("mysql", env.E3user+":"+env.E3pwd+"@tcp("+env.E3host+":"+env.E3port+")/"+env.E3db+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		logger.Log.Error(err)
	}
	env.RDB = database0

	model.InitModels(database0)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", start)

	r := e.Group("/")

	controler.APIConrolerBooks(r)
	controler.APIConrolerMember(r)
	controler.APIConrolerPublisher(r)

	e.Logger.Fatal(e.Start(":8080"))

}

func start(c echo.Context) error {
	return c.String(http.StatusOK, "Library System api")
}

package controler

import (
	"net/http"
	"new_library/api_echo/publisher"

	"github.com/labstack/echo/v4"
)

func AddPublisher(c echo.Context) error {
	publisher, err := publisher.AddPublisher(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, publisher)
	}

}

func GetPublisher(c echo.Context) error {
	publisher, err := publisher.GetMember(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, publisher)
	}

}

func APIConrolerPublisher(g *echo.Group) {
	g.POST("api/publisher/add", AddPublisher)
	g.GET("api/publisher/get", GetPublisher)

}

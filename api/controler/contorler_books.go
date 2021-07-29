package controler

import (
	"net/http"
	"new_library/api_echo/books"

	"github.com/labstack/echo/v4"
)

func AddBooks(c echo.Context) error {
	books, err := books.AddBooks(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, books)
	}

}

func GetBooks(c echo.Context) error {
	books, err := books.GetBooks(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, books)
	}

}

func APIConrolerBooks(g *echo.Group) {
	g.POST("api/books/add", AddBooks)
	g.GET("api/books/get", GetBooks)
}

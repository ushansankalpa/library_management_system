package books

import (
	"new_library/model"

	op "new_library/operations/books"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) ([]*model.Books, error) {
	result, err := op.GetBooks()
	return result, err
}

package books

import (
	"new_library/model"
	op "new_library/operations/books"

	"github.com/labstack/echo/v4"
)

func AddBooks(c echo.Context) (*model.Books, error) {
	books := model.Books{}
	if error := c.Bind(&books); error != nil {
		return nil, error
	}

	result, err := op.AddBooks(&books)
	return result, err
}

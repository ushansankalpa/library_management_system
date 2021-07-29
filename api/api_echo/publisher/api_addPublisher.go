package publisher

import (
	"new_library/model"
	op "new_library/operations/publisher"

	"github.com/labstack/echo/v4"
)

func AddPublisher(c echo.Context) (*model.Publishers, error) {
	publisher := model.Publishers{}
	if error := c.Bind(&publisher); error != nil {
		return nil, error
	}

	result, err := op.AddPublisher(&publisher)
	return result, err
}

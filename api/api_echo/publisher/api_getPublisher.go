package publisher

import (
	"new_library/model"
	op "new_library/operations/publisher"

	"github.com/labstack/echo/v4"
)

func GetMember(c echo.Context) ([]*model.Publishers, error) {
	result, err := op.GetPublisher()

	return result, err
}

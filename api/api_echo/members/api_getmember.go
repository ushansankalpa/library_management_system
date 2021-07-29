package members

import (
	"new_library/model"
	op "new_library/operations/members"

	"github.com/labstack/echo/v4"
)

func GetMember(c echo.Context) ([]*model.Members, error) {
	result, err := op.GetMember()

	return result, err
}

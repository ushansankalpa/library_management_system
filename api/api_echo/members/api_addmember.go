package members

import (
	"new_library/model"
	op "new_library/operations/members"

	"github.com/labstack/echo/v4"
)

func AddMember(c echo.Context) (*model.Members, error) {
	members := model.Members{}
	if error := c.Bind(&members); error != nil {
		return nil, error
	}

	result, err := op.AddBooks(&members)
	return result, err
}

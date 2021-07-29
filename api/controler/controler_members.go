package controler

import (
	"net/http"
	"new_library/api_echo/members"

	"github.com/labstack/echo/v4"
)

func AddMember(c echo.Context) error {
	members, err := members.AddMember(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, members)
	}

}

func GetMember(c echo.Context) error {
	members, err := members.GetMember(c)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	} else {
		return c.JSON(http.StatusOK, members)
	}

}

func APIConrolerMember(g *echo.Group) {
	g.POST("api/member/add", AddMember)
	g.GET("api/member/get", GetMember)

}

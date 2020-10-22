package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/skanehira/golang-sample-api/service"
)

type User struct {
	su *service.User
}

func NewUserHandler(su *service.User) *User {
	return &User{su: su}
}

func (u *User) Users(c echo.Context) error {
	users, err := u.su.Users()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

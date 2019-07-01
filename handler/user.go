package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type UserHandler struct {
	repo repository.UserRepo
}


func (u UserHandler) GetList(c echo.Context) (err error) {
	var users []model.User
	err = u.repo.GetListUsers(&users)
	if err != nil {
		return RespondToClient(c, 404, "Get users fail", nil)
	}
	return RespondToClient(c, 200, "Get users success", users)
}
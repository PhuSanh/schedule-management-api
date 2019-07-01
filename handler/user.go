package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
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

func (u UserHandler) Update(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	user := new(model.User)
	if err = c.Bind(&user); err != nil {
		return RespondToClient(c, 200, "Login request params wrong", err)
	}
	user.ID = id
	err = u.repo.UpdateUser(user)

	if err != nil {
		return RespondToClient(c, 404, "Get users fail", nil)
	}
	return RespondToClient(c, 200, "Get users success", user)
}

func (u UserHandler) Delete(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	err = u.repo.DeleteUser(id)
	if err != nil {
		return RespondToClient(c, 404, "Get users fail", nil)
	}
	return RespondToClient(c, 200, "Get users success", true)
}
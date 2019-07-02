package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type UserCategoryHandler struct {
	repo repository.UserCategoryRepo
}

func (ch UserCategoryHandler) GetList(c echo.Context) (err error) {
	var listUserCategories []model.UserCategory
	err = ch.repo.GetListUserCategories(&listUserCategories)
	if err != nil {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, listUserCategories)
}

func (ch UserCategoryHandler) Update(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	userCategory := new(model.UserCategory)
	if err = c.Bind(&userCategory); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	userCategory.ID = id
	err = ch.repo.UpdateUserCategory(userCategory)

	if err != nil {
		return RespondToClient(c, common.ERROR_UPDATE_ROW_IN_DB, common.MSG_UPDATE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, userCategory)
}

func (ch UserCategoryHandler) Delete(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	err = ch.repo.DeleteUserCategory(id)
	if err != nil {
		return RespondToClient(c, common.ERROR_DELETE_ROW_IN_DB, common.MSG_DELETE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, true)
}
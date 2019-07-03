package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type UserGroupHandler struct {
	repo repository.GroupRepo
}

func (gh UserGroupHandler) GetList(c echo.Context) (err error) {
	var listGroups []model.Group
	err = gh.repo.GetListGroups(&listGroups)
	if err != nil {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, listGroups)
}

func (gh UserGroupHandler) Create(c echo.Context) (err error) {
	group := new(model.Group)
	if err = c.Bind(&group); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	err = gh.repo.CreateGroup(group)
	if err != nil {
		return RespondToClient(c, common.ERROR_INSERT_ROW_TO_DB, common.MSG_INSERT_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, group)
}

func (gh UserGroupHandler) Update(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	group := new(model.Group)
	if err = c.Bind(&group); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	group.ID = id
	err = gh.repo.UpdateGroup(group)

	if err != nil {
		return RespondToClient(c, common.ERROR_UPDATE_ROW_IN_DB, common.MSG_UPDATE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, group)
}

func (gh UserGroupHandler) Delete(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	err = gh.repo.DeleteGroup(id)
	if err != nil {
		return RespondToClient(c, common.ERROR_DELETE_ROW_IN_DB, common.MSG_DELETE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, true)
}
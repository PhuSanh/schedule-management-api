package handler

import (
	"github.com/labstack/echo/v4"
	"schedule-management-api/common"
	"schedule-management-api/model"
	"schedule-management-api/repository"
)

type ScheduleHandler struct {
	repo repository.ScheduleRepo
}

func (sh ScheduleHandler) GetList(c echo.Context) (err error) {
	var listSchedules []model.Schedule
	err = sh.repo.GetListSchedules(&listSchedules)
	if err != nil {
		return RespondToClient(c, common.ERROR_GET_ROW_FROM_DB, common.MSG_GET_ROW_FROM_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, listSchedules)
}

func (sh ScheduleHandler) Create(c echo.Context) (err error) {
	schedule := new(model.Schedule)
	if err = c.Bind(&schedule); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	err = sh.repo.CreateSchedule(schedule)
	if err != nil {
		return RespondToClient(c, common.ERROR_INSERT_ROW_TO_DB, common.MSG_INSERT_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, schedule)
}

func (sh ScheduleHandler) Update(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	schedule := new(model.Schedule)
	if err = c.Bind(&schedule); err != nil {
		return RespondToClient(c, common.ERROR_REQUEST_DATA_INVALID, common.MSG_REQUEST_DATA_INVALID, err)
	}
	schedule.ID = id
	err = sh.repo.UpdateSchedule(schedule)

	if err != nil {
		return RespondToClient(c, common.ERROR_UPDATE_ROW_IN_DB, common.MSG_UPDATE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, schedule)
}

func (sh ScheduleHandler) Delete(c echo.Context) (err error) {
	id := common.ParseParamID(c.Param("id"))
	err = sh.repo.DeleteSchedule(id)
	if err != nil {
		return RespondToClient(c, common.ERROR_DELETE_ROW_IN_DB, common.MSG_DELETE_ROW_TO_DB, err)
	}
	return RespondToClient(c, common.ERROR_NO_ERORR, common.MSG_SUCEESS, true)
}
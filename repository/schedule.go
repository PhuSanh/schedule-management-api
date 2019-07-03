package repository

import (
	"schedule-management-api/database"
	"schedule-management-api/model"
)

type ScheduleRepo struct {}

func (sr ScheduleRepo) GetListSchedules(listSchedules *[]model.Schedule) (err error) {
	err = database.MysqlConn.Find(&listSchedules).Error
	return
}

func (sr ScheduleRepo) UpdateSchedule(schedule *model.Schedule) (err error) {
	err = database.MysqlConn.Model(&schedule).Updates(schedule).Error
	return
}

func (sr ScheduleRepo) CreateSchedule(schedule *model.Schedule) (err error) {
	err = database.MysqlConn.Create(&schedule).Error
	return
}

func (sr ScheduleRepo) DeleteSchedule(id int64) (err error) {
	err = database.MysqlConn.Where("id = ?", id).Delete(&model.Schedule{}).Error
	return
}
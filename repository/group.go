package repository

import (
	"schedule-management-api/database"
	"schedule-management-api/model"
)

type GroupRepo struct {}

func (gr GroupRepo) GetListGroups(listGroups *[]model.Group) (err error) {
	err = database.MysqlConn.Find(&listGroups).Error
	return
}

func (gr GroupRepo) UpdateGroup(group *model.Group) (err error) {
	err = database.MysqlConn.Model(&group).Updates(group).Error
	return
}

func (gr GroupRepo) CreateGroup(group *model.Group) (err error) {
	err = database.MysqlConn.Create(&group).Error
	return
}

func (gr GroupRepo) DeleteGroup(id int64) (err error) {
	err = database.MysqlConn.Where("id = ?", id).Delete(&model.Group{}).Error
	return
}
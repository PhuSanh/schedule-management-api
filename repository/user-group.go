package repository

import (
	"schedule-management-api/database"
	"schedule-management-api/model"
)

type UserGroupRepo struct {}

func (gr UserGroupRepo) GetListUsersInGroup(listUsers *[]model.UserGroup, groupId int64) (err error) {
	err = database.MysqlConn.Where(&model.UserGroup{GroupId: groupId}).Find(&listUsers).Error
	return
}

func (gr UserGroupRepo) AddUserInGroup(userGroup *model.UserGroup) (err error) {
	err = database.MysqlConn.Create(&userGroup).Error
	return
}

func (gr UserGroupRepo) DeleteUserFromGroup(userId, groupId int64) (err error) {
	err = database.MysqlConn.Where(&model.UserGroup{
		GroupId: groupId,
		UserId: userId,
	}).Delete(&model.UserGroup{}).Error
	return
}
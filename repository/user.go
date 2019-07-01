package repository

import (
	"fmt"
	"schedule-management-api/database"
	"schedule-management-api/model"
)

type UserRepoInterface interface {
	GetListUser() ([]model.User, error)
}

type UserRepo struct {}

func (ur UserRepo) GetListUsers(listUser *[]model.User) (err error) {
	err = database.MysqlConn.Find(&listUser).Error
	return
}

func (ur UserRepo) GetUserLogin(user *model.User, body model.LoginForm) (err error) {
	err = database.MysqlConn.Where(&model.User{Username: body.Username}).First(&user).Error
	return
}

func (ur UserRepo) UpdateUser(user *model.User) (err error) {
	err = database.MysqlConn.Model(&user).Updates(user).Error
	return
}

func (ur UserRepo) CreateUser(user *model.User) (err error) {
	err = database.MysqlConn.Create(&user).Error
	return
}

func (ur UserRepo) DeleteUser(id int64) (err error) {
	err = database.MysqlConn.Where("id = ?", id).Delete(&model.User{}).Error
	return
}
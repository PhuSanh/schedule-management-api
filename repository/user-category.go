package repository

import (
	"schedule-management-api/database"
	"schedule-management-api/model"
)

type UserCategoryRepo struct {}

func (cr UserCategoryRepo) GetListUserCategories(listCategories *[]model.UserCategory) (err error) {
	err = database.MysqlConn.Find(&listCategories).Error
	return
}

func (cr UserCategoryRepo) UpdateUserCategory(userCategory *model.UserCategory) (err error) {
	err = database.MysqlConn.Model(&userCategory).Updates(userCategory).Error
	return
}

func (cr UserCategoryRepo) CreateUserCategory(userCategory *model.UserCategory) (err error) {
	err = database.MysqlConn.Create(&userCategory).Error
	return
}

func (cr UserCategoryRepo) DeleteUserCategory(id int64) (err error) {
	err = database.MysqlConn.Where("id = ?", id).Delete(&model.UserCategory{}).Error
	return
}
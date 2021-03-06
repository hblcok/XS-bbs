package dao

import (
	"errors"

	"gorm.io/gorm"
	"xs.bbs/internal/pkg/constant/e"
)

// Insert 新增用户
func (u *UserDao) Insert(user *UserModel) (err error) {
	return u.db.Create(&user).Error
}

// Delete 根据用户ID删除用户，软删除
func (u *UserDao) Delete(userID int64) bool {
	return u.db.Where("user_id = ?", userID).Delete(&UserModel{}).RowsAffected > 0
}

// Update 根据用户ID修改用户
func (u *UserDao) Update(user *UserModel) error {
	return u.db.Where("user_id = ?").Updates(&user).Error
}

// SelectByName 根据用户名查询用户
func (u *UserDao) GetUserByName(userName string) (*UserModel, error) {
	var user UserModel
	if err := u.db.Where("username = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// SelectByID 根据用户ID查询用户
func (u *UserDao) GetUserByID(userID int64) (*UserModel, error) {
	var (
		user UserModel
		err  error
	)
	if err = u.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.ErrUserNotExist
		}
		return nil, err
	}
	return &user, nil
}

// CheckUserByUserName 根据userName检查用户是否存在
func (u *UserDao) CheckUserByUserName(userName string) error {
	var count int64
	if err := u.db.Model(&UserModel{}).Where("username = ?", userName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return e.ErrUserExist
	}
	return nil
}

// CheckUserByEmail 通过email检查用户
func (u *UserDao) CheckUserByEmail(email string) error {
	var count int64
	if err := u.db.Model(&UserModel{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return e.ErrEmailExist
	}
	return nil
}

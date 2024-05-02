package gormdb

import (
	"karaoke/domain/model"
	"karaoke/domain/repository"
	"karaoke/utils"
)

func NewUser() repository.UserRepository {
	return userRepository{}
}

type userRepository struct {
}

func (repo userRepository) GetUser(userId string, password string) (*model.User, error) {

	var user []*model.User
	DB.Table("user").Where("user_id  = ? and password = ?", userId, password).Find(&user)
	if len(user) == 0 {
		return nil, nil
	}
	return user[0], nil
}
func (repo userRepository) AddUser(userId string, password string, email string) error {
	var err error
	user := &model.User{
		UserID:      userId,
		Password:    password,
		Email:       email,
		Status:      "0",
		CreatedTime: utils.GettimeNow(),
		UpdatedTime: utils.GettimeNow(),
	}
	err = DB.Table("user").Create(user).Error

	if err != nil {
		return err
	}
	return nil
}
func (r userRepository) CheckUserName(userId string, email string) ([]*model.GetUser, error) {
	var listUser []*model.GetUser

	if len(listUser) == 0 {
		return nil, nil
	}

	return listUser, nil
}

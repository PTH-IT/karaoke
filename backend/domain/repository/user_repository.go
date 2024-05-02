package repository

import (
	"karaoke/domain/model"
)

type UserRepository interface {
	GetUser(userId string, password string) (*model.User, error)
	AddUser(userId string, password string, email string) error
	CheckUserName(userId string, email string) ([]*model.GetUser, error)
}

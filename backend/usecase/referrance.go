package usecase

import (
	"fmt"

	"karaoke/domain/model"
	"karaoke/domain/repository"
	InforLog "karaoke/log/infor"
)

type Reference interface {
	GetUser(userId string, password string) (*model.User, error)
	AddUser(userId string, password string, email string) error
	CheckUserName(userId string, email string) ([]*model.GetUser, error)
}

func NewReferrance(
	userRepository repository.UserRepository,
) Reference {
	return reference{
		userRepository,
	}
}

type reference struct {
	userRepository repository.UserRepository
}

func (r reference) GetUser(userId string, password string) (*model.User, error) {

	user, err := r.userRepository.GetUser(userId, password)
	return user, err
}

func (r reference) AddUser(userId string, password string, email string) error {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.AddUser call"))
	err := r.userRepository.AddUser(userId, password, email)
	return err
}
func (r reference) CheckUserName(userId string, email string) ([]*model.GetUser, error) {
	InforLog.PrintLog(fmt.Sprintf("r.mongoRepository.GetUser call"))
	user, err := r.userRepository.CheckUserName(userId, email)
	return user, err
}

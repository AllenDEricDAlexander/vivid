package repository

import (
	"user/model"
	"user/repo/persistence"
)

type UserRepository interface {
	Create(user model.User) error
}

type UserRepositoryImpl struct {
	userDAOContext persistence.UserDAOContext
	redisDAO       persistence.RedisDAO
	mySQLDAO       persistence.MySQLDAO
}

func (u *UserRepositoryImpl) Create(user model.User) {

}

package persistence

import "user/model"

type UserBuilder interface {
	FromUser(user model.User) (UserDO, error)
	ToUser(user UserDO) (model.User, error)
}

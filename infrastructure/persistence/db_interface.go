package persistence

import "kitchenmaniaapi/domain/entity"

type DbInterface interface {
	NewUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User,error)
}

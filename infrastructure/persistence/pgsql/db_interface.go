package pgsql

import "kitchenmaniaapi/domain/entity"

type DbInterface interface {
	NewUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
	CreatePost(blog *entity.Blog) error
	TokenInBlacklist(token *string) bool
}
package postgresdb

import (
	"kitchenmaniaapi/domain/entity"
	"kitchenmaniaapi/entity"

	"github.com/gofrs/uuid"
)

// type UserRepository interface {
// 	NewUser(user *entity.User) error
// 	FindUserByEmail(email string) (*entity.User, error)
// 	CreatePost(blog entity.Blog) error
// 	TokenInBlacklist(token *string) bool
// 	DeletePost(blogID, userID string) error
// 	UpdatePost(blog entity.Blog) error
// 	GetAllPosts(user entity.User) ([]entity.Blog, error)
// }

type UserRepository interface {
	GetAll() ([]entity.User, error)
	GetByID(email string) (entity.User, error)
	Add(user entity.User) error
	Update(user entity.User) error
	Delete(email string) error
}

type BlacklistRepository interface{
	TokenInBlacklist(token *string) bool
}

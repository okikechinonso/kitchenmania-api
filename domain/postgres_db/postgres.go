package postgresdb

import (
	"kitchenmaniaapi/domain/entity"

	"github.com/gofrs/uuid"
)

type UserRepository interface {
	NewUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
	CreatePost(blog entity.Blog) error
	TokenInBlacklist(token *string) bool
	DeletePost(blogID, userID string) error
	UpdatePost(blog entity.Blog) error
	GetAllPosts(user entity.User) ([]entity.Blog, error)
}

type Repository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}

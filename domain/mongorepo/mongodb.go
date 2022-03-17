package mongorepo

import "kitchenmaniaapi/entity"

type BlogRepository interface {
	GetAll() ([]entity.Blog, error)
	GetByID(email string) (entity.Blog, error)
	Add(blog entity.Blog) error
	Update(blog entity.Blog) error
	Delete(email string) error
}

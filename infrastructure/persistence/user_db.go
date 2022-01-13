package persistence

import (
	"kitchenmaniaapi/domain/entity"
)

func (d *Database) NewUser(user *entity.User) error {
	result := d.DB.Create(user)
	return result.Error
}
func (d *Database) FindUserByEmail(email string) (*entity.User,error) {
	var user entity.User
	err := d.DB.Where("email = ?", email).First(&user).Error
	return &user,err
}

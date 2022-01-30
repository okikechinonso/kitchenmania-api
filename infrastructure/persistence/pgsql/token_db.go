package pgsql

import (
	"kitchenmaniaapi/domain/entity"
)

func (d *Database) TokenInBlacklist(token *string) bool {
	result := d.DB.Where("token = ?", *token).Find(&entity.Blacklist{})
	return result.Error != nil

}

func (d *Database) AddToBlackList(blacklist *entity.Blacklist) error {
	result := d.DB.Create(blacklist)
	return result.Error
}

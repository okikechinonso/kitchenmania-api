// DB blog methods
package pgsql

import "kitchenmaniaapi/domain/entity"

func (d *Database) CreatePost(blog *entity.Blog) error {
	result := d.PgDB.Create(blog)
	return result.Error
}

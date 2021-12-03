package persistence

import "gorm.io/gorm"

type Database struct{
	DB *gorm.DB
}
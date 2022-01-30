package pgsql

import (
	"fmt"
	"kitchenmaniaapi/domain/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_MODE")
	var dns string
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == ""{
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	} else {
		dns = databaseurl
	}

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database %v", err)
	}

	log.Println("connected to database")
	d.DB = db
}
func (d *Database) Migrate() {
	err := d.DB.AutoMigrate(&entity.User{}, &entity.Blacklist{}, &entity.Blog{})
	if err != nil {
		log.Printf("%s", err)
	}
}

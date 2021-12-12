package persistence

import (
	"fmt"
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
	
	if len(host)==0 && len(user)==0 && len(password)==0 && len(dbName)==0{
		log.Fatal("No database credentials")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil{
		fmt.Errorf("unable to connect to database",err)
	}
	d.DB = db
}

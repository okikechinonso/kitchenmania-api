package postgresdb

import (
	"context"
	"fmt"
	"kitchenmaniaapi/entity"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB  struct {
	PostgresClient *gorm.DB
}

func (pg *PostgresDB) Init() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_MODE")
	var dns string
	databaseurl := os.Getenv("DATABASE_URL")
	if databaseurl == "" {
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	} else {
		dns = databaseurl
	}

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect to database postgresDB %v", err)
	}

	log.Println("connected to databases")
	pg.PostgresClient = db
}

func (pg *PostgresDB) Migrate() {
	err := pg.PostgresClient.AutoMigrate(&entity.User{}, &entity.Blacklist{})
	if err != nil {
		log.Printf("%s", err)
	}
}


func (mg *MongoDB) GetAll() ([]aggregate.Product, error)            { return nil, nil }
func (mg *MongoDB) GetByID(id uuid.UUID) (aggregate.Product, error) { return nil,nil }
func (mg *MongoDB) Add(product aggregate.Product) error             { return nil }
func (mg *MongoDB) Update(product aggregate.Product) error          { return nil }
func (mg *MongoDB) Delete(id uuid.UUID) error                       { return nil }
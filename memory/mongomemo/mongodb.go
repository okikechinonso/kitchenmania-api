package mongodb

import (
	"context"
	"log"
	"os"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	MongoClient *mongo.Client
}

func (mg *MongoDB) Init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("unable to connect to database MongoDB %v", err)
	}
	mg.MongoClient = client
}

func (mg *MongoDB) GetAll() ([]aggregate.Product, error)            { return nil, nil }
func (mg *MongoDB) GetByID(id uuid.UUID) (aggregate.Product, error) { return nil,nil }
func (mg *MongoDB) Add(product aggregate.Product) error             { return nil }
func (mg *MongoDB) Update(product aggregate.Product) error          { return nil }
func (mg *MongoDB) Delete(id uuid.UUID) error                       { return nil }

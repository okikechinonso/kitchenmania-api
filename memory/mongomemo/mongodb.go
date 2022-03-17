package mongodb

import (
	"context"
	"kitchenmaniaapi/entity"
	"log"
	"os"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	MongoClient *mongo.Client
}

func (mg *MongoDB) Init() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatalf("unable to connect to MongoDB MongoDB %v", err)
	}
	mg.MongoClient = client
}

func (mg *MongoDB) GetAll(user entity.User) ([]entity.Blog, error) {
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	filter := bson.M{"user_id": user.ID}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	var result []entity.Blog
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Println(err)
	}
	return result, err
 }


func (mg *MongoDB) GetByID(id uuid.UUID) (entity.Blog, error) { return entity.Blog{}, nil }

func (mg *MongoDB) Add(blog entity.Blog) error {
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_, err := coll.InsertOne(context.TODO(), blog)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (mg *MongoDB) Update(blog entity.Blog) error {
	coll := mg.MongoClient.Database("kitchenmania").Collection("blog")
	filter := bson.M{"user_id": blog.UserID, "author": blog.Author}
	log.Println(blog.Author, blog.UserID)
	err := coll.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": blog}).Decode(&entity.Blog{})
	if err != nil {
		log.Println(err)
	}
	return err
}
func (mg *MongoDB) Delete(blogID, userID string) error { 
	coll := mg.MongoClient.Database("kitchenmania").Collection("blog")
	filter := bson.M{"user_id": userID, "blog_id": blogID}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	return err
	 }



func (mg *MongoDB) GetAllPosts(user entity.User) ([]entity.Blog, error) {
	coll := mg.MongoClient.Database("kitchenmania").Collection("blog")
	filter := bson.M{"user_id": user.ID}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Println(err)
	}
	var result []entity.Blog
	err = cursor.All(context.TODO(), &result)
	if err != nil {
		log.Println(err)
	}
	return result, err
}




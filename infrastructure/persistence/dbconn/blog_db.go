package dbconn

import (
	"context"
	"kitchenmaniaapi/domain/entity"
	"log"

	"github.com/globalsign/mgo/bson"
)

func (d *Database) CreatePost(blog entity.Blog) error {
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_, err := coll.InsertOne(context.TODO(), blog)
	if err != nil{
		log.Println(err)
	}
	return err

}

func (d *Database) DeletePost(blogID, userID string)error{
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_,err := coll.DeleteOne(context.TODO(), blogID)
	if err != nil {
		log.Println(err)
	}
	return err
}

func(d *Database) UpdatePost(blog entity.Blog) error{
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	filter := bson.M{"user_id":blog.UserID,"author":blog.Author}
	log.Println(blog.Author,blog.UserID)
	// update := bson.D{{"$set",blog}}
	err := coll.FindOneAndUpdate(context.TODO(),filter, bson.M{"$set":blog}).Decode(&entity.Blog{})
	if err != nil {
		log.Println(err)
	}
	return err
}

func(d *Database) ViewPosts(userID string) ([]entity.Blog,error){
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_,err := coll.Find(context.TODO(), userID)
	if err != nil {
		log.Println(err)
	}
	return []entity.Blog{},err
}
func(d *Database) GetAllPosts() ([]entity.Blog,error){
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_,err := coll.Find(context.TODO(), "")
	if err != nil {
		log.Println(err)
	}
	return []entity.Blog{},err
}



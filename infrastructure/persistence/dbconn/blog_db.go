package dbconn

import (
	"context"
	"kitchenmaniaapi/domain/entity"
	"log"
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

func(d *Database) UpdatePost(blogID string) error{
	coll := d.MongoClient.Database("kitchenmania").Collection("blog")
	_,err := coll.UpdateByID(context.TODO(),blogID, entity.Blog{})
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



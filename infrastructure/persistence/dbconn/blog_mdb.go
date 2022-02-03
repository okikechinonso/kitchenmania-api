package dbconn

import (
	"context"
	"kitchenmaniaapi/domain/entity"
	"log"
)

func (d *Database) CreatePost(blog entity.Blog) error {
	coll := d.MongoClient.Database("kitchen").Collection("blog")
	_, err := coll.InsertOne(context.TODO(), blog)
	if err != nil{
		log.Println(err)
	}
	return err

}

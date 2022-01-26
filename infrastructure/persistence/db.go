package persistence

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
	
	dns := ""
	if len(host)==0 && len(user)==0 && len(password)==0 && len(dbName)==0{
		dns = "postgres://fvaqghjvohtjwq:523685305314609ce4b5d92e1eb2ecc3416abc2915fbdac25bfe871df8b1f6f8@ec2-52-31-219-113.eu-west-1.compute.amazonaws.com:5432/d4o19h76ppqjoo"
		log.Fatal("No database credentials")
	}else{
		dns = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbName, port, sslmode)
	}
	
	
	db, err := gorm.Open(postgres.Open(dns),&gorm.Config{})
	if err != nil{
		log.Fatalf("unable to connect to database %v",err)
	}

	log.Println("connected to database")
	d.DB = db
}
func (d *Database) Migrate(){
	err := d.DB.AutoMigrate(&entity.User{},&entity.Blacklist{},&entity.Blog{})
	if err != nil {
		log.Printf("%s",err)
	}
}

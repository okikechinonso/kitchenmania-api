package entity


type Blog struct {
	Model
	UserID   string    `json:"user_id" bson:"user_id" gorm:"foreignkey:User(id)"`
	Author   string    `json:"author" bson:"author"` 
	Title    string    `json:"title" gorm:"not null"  bson:"title" binding:"required" form:"title"`
	Body     string    `json:"body" gorm:"not null" bson:"body" binding:"required" form:"body"`
	Images   []Image   `gorm:"many2many:blog_images" bson:"images" json:"images" binding:"required"`
	Comments []Comment `json:"comments,omitempty" bson:"comments,omitempty"`
	Like     int64     `json:"like,omitempty"`
}

type Comment struct {
	Model
	BlogID   string `json:"blog_id" bson:"blog_id" binding:"required" gorm:"foreignkey:Blog(id)"`
	UserID   string `json:"user_id" bson:"user_id" binding:"required" gorm:"foreignkey:User(id)"`
	UserName string `json:"user_name" bson:"user_name" binding:"required" gorm:"not null"`
	Body     string `json:"body" bson:"body" gorm:"not null" binding:"required" form:"body"`
}

type Image struct {
	Model
	BlogID   string `gorm:"foreignkey:Blog(id)" binding:"required"`
	ImageUrl string `gorm:"not null" binding:"reqiured"`
}

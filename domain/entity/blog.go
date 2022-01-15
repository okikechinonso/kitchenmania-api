package entity

type Blog struct {
	Model
	UserID   string    `json:"user_id" gorm:"foreignkey:User(id)"`
	Author   string    `json:"author"`
	Title    string    `json:"title" gorm:"not null" binding:"required,title" form:"title"`
	Body     string    `json:"body" gorm:"not null" binding:"required,body" form:"body"`
	Images   []string  `json:"images" gorm:"not null" binding:"required,images" form:"images"`
	Comments []Comment `json:"comments,omitempty"`
	Like     int64     `json:"like,omitempty"`
}

type Comment struct {
	Model
	UserID   string	`json:"user_id" gorm:"foreignkey:User(id)"`
	UserName string	`json:"user_name" gorm:"not null"`
	Body     string `json:"body" gorm:"not null" binding:"required,body" form:"body"`
}

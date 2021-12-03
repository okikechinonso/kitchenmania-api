package entity

import "time"

type Model struct {
	ID        string `sql:"type:uuid; default:uuid_generate_v4();size:100; not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
type User struct {
	Model
	FirstName  string     `json:"first_name" gorm:"not null" binding:"required" form:"first_name"`
	LastName   string     `json:"last_name" gorm:"not null" binding:"required" form:"last_name"`
	Email      string     `json:"email" gorm:"not null" binding:"required,email" form:"email"`
	Password   string     `json:"password" gorm:"-" binding:"required" form:"password"`
	Phone      string     `json:"phone,omitempty"`
	Bio        string     `json:"bio,omitempty"`
	Followers  []Follower `gorm:"many2many:followers" json:"followers,omitempty"`
	Verified   bool       `json:"verified" gorm:"default:false"`
	Active     bool       `json:"active" gorm:"default:false"`
	Blogs      []Blog     `gorm:"-" json:"blogs,omitempty"`
	ImageUrl   string     `json:"image_url,omitempty"`
	Profession string     `json:"profession,omitempty"`
}

type Follower struct {
	User
}

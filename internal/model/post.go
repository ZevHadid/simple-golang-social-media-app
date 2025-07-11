package model

import "time"

type Post struct {
	ID       uint `gorm:"primaryKey"`
	UserID   uint
	Title    string
	Content  string
	PostDate time.Time
	User     User `gorm:"foreignKey:UserID"`
}

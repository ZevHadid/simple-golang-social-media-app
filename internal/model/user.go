package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

package model

type User struct {
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
}

package model

type User struct {
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
}

type Claims struct {
	Username string `json:"username"`
}

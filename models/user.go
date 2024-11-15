package models

type User struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Email    string
	Password string
}

type RegisterUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
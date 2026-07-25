package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"uniqueIndex`
	Password string 
	UserName string 
}

func NewUser (email string, password string, userName string) *User{
	return &User{
		Email: email,
		Password: password,
		UserName: userName,
	}
}

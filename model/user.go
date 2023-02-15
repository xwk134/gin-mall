package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"type:varchar(20);not null;unique"`
	Email          string `gorm:"type:varchar(20);not null"`
	PasswordDigest string `gorm:"type:varchar(100);not null"`
	NickName       string `gorm:"type:varchar(20);not null"`
	Status         string `gorm:"type:varchar(20)"`
	Avatar         string `gorm:"type:varchar(20)"`
	Money          string `gorm:"type:varchar(20)"`
}

const (
	PasswordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

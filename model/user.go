package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"type:varchar(20);not null;unique"`
	Email          string `gorm:"type:varchar(20);not null"`
	PasswordDigest string `gorm:"type:varchar(20);not null"`
	NickName       string `gorm:"type:varchar(20);not null"`
	Status         string `gorm:"type:varchar(20);not null"`
	Avatar         string `gorm:"type:varchar(20);not null"`
	Money          string `gorm:"type:varchar(20);not null"`
}

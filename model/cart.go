package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	BossId    uint `gorm:"not null comment:商家id"`
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"`
	Check     bool `gorm:"comment:是否支付"`
}

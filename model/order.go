package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"type:bigint(20);not null"`
	ProductId uint `gorm:"type:bigint(20);not null"`
	BossId    uint `gorm:"type:bigint(20);not null"`
	AddressId uint `gorm:"type:bigint(20);not null"`
	Num       int
	OrderNum  uint64
	Type      uint `gorm:"type:bigint(20);comment:1 未支付 2 已支付"`
	Money     float64
}

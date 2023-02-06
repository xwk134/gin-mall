package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20);not null"`
	Category      uint   `gorm:"type:bigint(20);not null"`
	Title         string `gorm:"type:varchar(50);not null"`
	Info          string `gorm:"type:varchar(200);not null"`
	ImgPath       string `gorm:"type:varchar(256);not null"`
	Price         string `gorm:"type:varchar(20);not null"`
	DiscountPrice string `gorm:"type:varchar(20);not null"`
	OnSale        bool   `gorm:"default:false"`
	Num           int    `gorm:"type:bigint(20);not null"`
	BossId        uint   `gorm:"type:bigint(20);not null"`
	BossName      string `gorm:"type:varchar(20);not null"`
	BossAvatar    string `gorm:"type:varchar(20);not null"`
}

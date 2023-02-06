package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	User      User    `gorm:"ForeignKey:UserId"`
	UserId    uint    `gorm:"type:bigint(20);not null"`
	Product   Product `gorm:"ForeignKey:ProductId"`
	ProductId uint    `gorm:"type:bigint(20);not null"`
	Boss      User    `gorm:"ForeignKey:BossId"`
	BossId    uint    `gorm:"type:bigint(20);not null"`
}

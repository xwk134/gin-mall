package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID  uint   `gorm:"not null;comment:用户id"`
	Name    string `gorm:"type:varchar(20);not null;comment:收货人姓名"`
	Phone   string `gorm:"type:varchar(11);not null;comment:收货人手机号"`
	Address string `gorm:"type:varchar(50);not null;comment:收货地址"`
}

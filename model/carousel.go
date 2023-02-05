package model

import "gorm.io/gorm"

type Carousel struct {
	gorm.Model
	ImgPath   string `gorm:"not null comment:轮播图地址"`
	ProductId uint   `gorm:"not null comment:商品id"`
}

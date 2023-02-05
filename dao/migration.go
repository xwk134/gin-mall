package dao

import (
	"fmt"
	"gin-mall/model"
)

func migration() {
	err := _db.Set("grom:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Address{},
			&model.Admin{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.Product{},
			&model.ProductImg{},
			&model.User{},
		)
	if err != nil {
		fmt.Println("err", err)
	}
	return
}

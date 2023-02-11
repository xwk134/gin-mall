package routes

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//跨域
	r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "suucess")

		})
		//用户操作
		v1.POST("user/register", api.UserRegister)
	}
	return r
}

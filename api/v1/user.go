package v1

import (
	"gin-mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var UserRegister service.UserService
	if err := c.ShouldBind(&UserRegister); err == nil {
		res := UserRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

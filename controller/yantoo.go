package controller

import (
	"gin-template/common"
	"github.com/gin-gonic/gin"
)

// Hello 控制器函数
// @Summary Hello endpoint
// @Description Returns a greeting message
// @ID hello
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response{data=string}
// @Router /hello [get]
func Hello(c *gin.Context) {
	c.JSON(200, common.Success("你好，登录成功"))
}

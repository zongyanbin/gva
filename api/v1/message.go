package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateMessage(c *gin.Context)  {
	var message model.Message
	_ = c.ShouldBindJSON(&message)
	fmt.Println(message)
	if err := service.CreateMessage(message); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}
func CreateMessageaa(c *gin.Context)  {
	var message model.Message
	message.Title = c.Request.FormValue("title")
	message.Content = c.Request.FormValue("content")
	err := c.ShouldBindJSON(&message)

	if(err != nil){
		c.JSON(200, gin.H{
			"msg":"成功了",
			"data":message,
		})
	}else {
		c.JSON(200, gin.H{
			"msg":"成功了",
			"data":message,
		})
	}
}


package router

import (
	"gin-vue-admin/api/app"

	"github.com/gin-gonic/gin"
)

func InitAppRouter(r *gin.RouterGroup) {
	q := r.Group("question")
	{
		q.GET("/", app.GetQuestionList) // 合并文件
	}
}

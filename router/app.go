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
	PaperRouter :=r.Group("paper")
	{
		PaperRouter.GET("/",app.GetPaperList) //获取试卷ID相关信列表
	}
}

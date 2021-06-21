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

	// 图片上传 扩展 extend
	upload := r.Group("extend")
	{
		upload.POST("/upload",app.UploadFile)                                 // 上传文件
	}

	Web_user_paper_answerRouter := r.Group("answer_question")
	{
		Web_user_paper_answerRouter.POST("createAnswer_question_front", app.CreateUser_paper_answer)
	}

}


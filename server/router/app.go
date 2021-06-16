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

	// 图片上传
	upload := r.Group("upload")
	{
		upload.POST("/",app.UploadFile)                                 // 上传文件
		//FileUploadAndDownloadGroup.POST("/getFileList", v1.GetFileList)                           // 获取上传文件列表
		//FileUploadAndDownloadGroup.POST("/deleteFile", v1.DeleteFile)                             // 删除指定文件
		//FileUploadAndDownloadGroup.POST("/breakpointContinue", v1.BreakpointContinue)             // 断点续传
		//FileUploadAndDownloadGroup.GET("/findFile", v1.FindFile)                                  // 查询当前文件成功的切片
		//FileUploadAndDownloadGroup.POST("/breakpointContinueFinish", v1.BreakpointContinueFinish) // 查询当前文件成功的切片
		//FileUploadAndDownloadGroup.POST("/removeChunk", v1.RemoveChunk)
	}
	//QuestionRouter.POST("createQuestion", v1.CreateQuestion)   // 新建Question
	//QuestionRouter.DELETE("deleteQuestion", v1.DeleteQuestion) // 删除Question
	//QuestionRouter.DELETE("deleteQuestionByIds", v1.DeleteQuestionByIds) // 批量删除Question
	//QuestionRouter.PUT("updateQuestion", v1.UpdateQuestion)    // 更新Question
	//QuestionRouter.GET("findQuestion", v1.FindQuestion)        // 根据ID获取Question
	//QuestionRouter.GET("getQuestionList", v1.GetQuestionList)  // 获取Question列表

}

package app

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// http://localhost:8888/app/question/?qid=1
type Request_paper struct {
	Paper_id int `form:"paperid"`
}

// 获取试卷信息
func GetPaperList(c *gin.Context)  {
	var requestPaper Request_paper
	c.ShouldBindQuery(&requestPaper)

	query := global.GVA_DB.Model(&model.Exam_paper{})
	if requestPaper.Paper_id > 0 {
		query.Where("id", requestPaper.Paper_id)
	}

	paper := []model.Exam_paper{}
	// 查找试卷和相关联的題
	err :=query.Preload("Question").Limit(10).Find(&paper).Error
	  for _,v := range paper{
		fmt.Println("v.Question",v.Question)

		  query := global.GVA_DB.Model(&model.Question{})
		  if v.ID > 0 {
			  query.Where("id", v.ID)
		  }
		  que := []model.Question{}
		  // 查处题目和其关联的选项
		  err := query.Preload("Question_options").Limit(10).Find(&que).Error
		  if err != nil {
			  global.GVA_LOG.Error("error!", zap.Any("err", err))
			  response.FailWithMessage("error", c)
		  }

		  // Question_options 想把que值付給 Question
		  v.Question = que
		  fmt.Println("Question_options:",que)
	}

	if err !=nil {
		global.GVA_LOG.Error("error",zap.Any("err",err))
			response.FailWithMessage("error",c)
	}else {
		response.OkWithDetailed(paper, "success", c)
	}
}

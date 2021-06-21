package app

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Request_paper struct {
	Paper_id int `form:"paperid"`
}

// 获取试卷信息
func GetPaperList(c *gin.Context){
	var requestPaper Request_paper
	c.ShouldBindQuery(&requestPaper)
	// service 服务层
	fmt.Println(requestPaper.Paper_id)
	if err,list_paper := service.GetQuestList(requestPaper.Paper_id); err != nil{
		global.GVA_LOG.Error("暂无数据",zap.Any("err", err))
		response.FailWithMessage("暂无数据",c)
	}else{
		if(list_paper.ID==0){
			response.OkWithMessage("没有创建试卷",c)
		}else{
			response.OkWithData(gin.H{"list_paper": list_paper},c)
		}
	}
}

// 获取试卷信息old
func GetPaperListTEST(c *gin.Context)  {
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

	//// 查出未过期的出差活动
	//func GetAllBusiness(projectid int64) (business []Business, err error) {
	//	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//	//join一定要select,其他不用select的话默认查询全部。
	//	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	//	db := GetDB()
	//	err = db.Order("business.updated_at desc").
	//		Preload("BusinessUsers").Preload("BusinessUsers.NickNames").Where("business.project_id = ?", projectid).
	//		Where("end_date > ?", time.Now()).
	//		Find(&business).Error
	//	return business, err
	//}

}

// @Tags User_paper_answer
// @Summary 创建User_paper_answer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "创建User_paper_answer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/createUser_paper_answer [post]
func CreateUser_paper_answer(c *gin.Context) {
	var User_paper_answer[] model.User_paper_answer
	_ = c.ShouldBindJSON(&User_paper_answer)

	fmt.Println("User_paper_answer",User_paper_answer)
	if err := service.CreateUser_paper_answer(User_paper_answer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func GetQuestList1(c *gin.Context)  {
	// prelad里面不是对应的表名字,而只主表中字段名字
	// join 一定要select,其它不用select的话默认查询全部
	// Prelad("Question.Question_options") 一嵌套加载！！
}
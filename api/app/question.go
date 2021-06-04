package app

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)
// http://localhost:8888/app/question/?QID=1
type Request struct {
	QID int
}

func GetQuestionList(c *gin.Context) {
	var req Request
	c.ShouldBindQuery(&req)
	query := global.GVA_DB.Model(&model.Question{})
	if req.QID > 0 {
		query.Where("id", req.QID)
	}
	que := []model.Question{}
	err := query.Limit(10).Find(&que).Error
	if err != nil {
		global.GVA_LOG.Error("error!", zap.Any("err", err))
		response.FailWithMessage("error", c)
	} else {
		response.OkWithDetailed(que, "success", c)
	}
}

package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitWx_userRouter(Router *gin.RouterGroup) {
	Wx_userRouter := Router.Group("wx_user").Use(middleware.OperationRecord())
	{
		Wx_userRouter.POST("createWx_user", v1.CreateWx_user)   // 新建Wx_user
		Wx_userRouter.DELETE("deleteWx_user", v1.DeleteWx_user) // 删除Wx_user
		Wx_userRouter.DELETE("deleteWx_userByIds", v1.DeleteWx_userByIds) // 批量删除Wx_user
		Wx_userRouter.PUT("updateWx_user", v1.UpdateWx_user)    // 更新Wx_user
		Wx_userRouter.GET("findWx_user", v1.FindWx_user)        // 根据ID获取Wx_user
		Wx_userRouter.GET("getWx_userList", v1.GetWx_userList)  // 获取Wx_user列表
	}
}

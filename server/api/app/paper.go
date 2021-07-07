package app

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/res"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Request_paper struct {
	Paper_id int `form:"paperid"`
}

// @Tags GetPaperList
// @Summary GetPaperList
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Request_paper false "试卷ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /paper/GetPaperList [GET]
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

// @Tags WXLogin
// @Summary 微信登录 WXLogin
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.User_paper_answer true "微信登录WXLogin"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /User_paper_answer/createUser_paper_answer [post]


func  WXLogin(code string) (*response.WXloginResp,error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	//合成url, 这里的appid和secret是在微信公众平台上获取的
	url = fmt.Sprintf(url, appID, appSecret, code)

	fmt.Println(url)
	// 创建http get 请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer  resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := response.WXloginResp{}
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode !=0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	fmt.Println(&wxResp)
	return  &wxResp, nil
}

// wechat/applet_login ?code=xxx [get] 路由
// 微信小程序登录

func AppletWeChatLogin( c *gin.Context) {
	code := c.Query("code")	// 获取code
	fmt.Println(code)
	// 根据code获取 openID和session_key
	wxLoginResp, err := service.WXLogin(code)
	if err != nil {
		res.SendResponse(c, nil, err.Error())
		return
	}
	// 初始化session对象 保存登录状态 	fmt.Println(wxLoginResp.SessionKey)  // openid oGFuX5IekCB64Z3i74HS2uJYjBWA session id

	session := sessions.Default(c)
	session.Set("openid", wxLoginResp.OpenId)
	session.Set("sessionKey", wxLoginResp.SessionKey)
	session.Save()

	//	$tore := sessions.NewSession()
	//var store = sessions.NewCookieStore([]byte("secret"))

	// 这里可以用openid和sessionkey的串接,或者使用你自己的规则进行拼接,然后进行MD5之后作为该用户的自定义登录态， 要保证mySession唯一,
	mySession := GetMD5Encode(wxLoginResp.OpenId + wxLoginResp.SessionKey)
	//
	//// 接下来可以将openid和sessionkey, mySession 存储到数据库中
	//// 但这里要保证 mySession 唯一， 以便用mySession去索引openid 和 sessionkey
	//c.String(200, wxLoginResp)
	//c.JSON(200,gin.H{
	//	"msg":"baocuo",
	//	"data":wxLoginResp,
	//	//"mySession":mySession,
	//})
	// 接下来可以将openid 和 sessionkey, mySession 存储到数据库或缓存中, 可以用mySession去索引openid 和sessionkey

	res.SendResponse(c, nil, mySession)
}


// 将一个字符串进行MD5加密后返回加密后的字符串
func GetMD5Encode(data string) string  {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

//校验微信返回的用户数据
func ValidateUserInfo(rawData, sessionKey, signature string) bool {
	signature2 := GetSha1(rawData + sessionKey)
	return signature == signature2
}
// SHA-1 加密
func GetSha1(str string) string {
	data := []byte(str)
	has := sha1.Sum(data)
	res := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return res
}




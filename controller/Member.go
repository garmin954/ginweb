package controller

import (
	"ginweb/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (mbc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mbc.sendSmsCode)
}

func (mbc *MemberController) sendSmsCode(context *gin.Context) {

	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败！",
		})
		return
	}

	mbs := service.MemberService{}
	isSend := mbs.SendSmsCode(phone)

	if isSend {
		context.JSON(http.StatusOK, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功！",
		})
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败！",
	})
}

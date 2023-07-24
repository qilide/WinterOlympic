package routes

import (
	"WinterOlympic/controller/account"
	"github.com/gin-gonic/gin"
)

//注册账户功能模块路由

func AccountGroupRoute(AccountGroup *gin.RouterGroup) {
	//用户登录
	AccountGroup.POST("/login", account.UserLogin)
	//用户注册
	AccountGroup.POST("/register", account.UserRegister)
	//获得验证码
	AccountGroup.POST("/mail", account.GetMail)
	//用户注销
	AccountGroup.GET("/logout", account.UserLogout)
}

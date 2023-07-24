package account

import (
	"WinterOlympic/controller/response"
	"WinterOlympic/dao/redis"
	"WinterOlympic/logic/account"
	"WinterOlympic/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserLogin 用户登录
// @Summary 用户登录
// @Description 用于用户登录
// @Tags 登录
// @Accept application/json
// @Produce application/json
// @Param object body LoginBinder true "登录参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.SuccessLogin "登陆成功"
// @failure 401 {object}  response.Information "邮箱或者密码错误"
// @failure 402 {object}  response.Information "请输入邮箱或者密码"
// @failure 403 {object}  response.Information "该用户未注册"
// @Router /account/login [POST]
func UserLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if len(email) == 0 || len(password) == 0 {
		var reb LoginBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.Json(c, 403, "请输入完整的信息", 0)
			return
		} else {
			email = reb.Email
			password = reb.Password
		}
	}
	var la account.Account
	user, loginErr := la.Login(email, password)
	if loginErr == account.PasswordErr {
		response.Json(c, 401, "邮箱或者密码错误", 0)
		return
	} else if loginErr == account.NotRegister {
		response.Json(c, 403, "该用户未注册", 0)
		return
	} else {
		token, _ := jwt.GenToken(user.ID, user.Email)
		var tr redis.TokenRedis
		tr.SetToken(user.Email, token)
		response.LoginJson(c, http.StatusOK, "登陆成功", user, token)
		return
	}
}

// UserRegister 新用户注册
// @Summary 新用户注册
// @Description 用于新用户注册账号使用
// @Tags 注册
// @Accept application/json
// @Produce application/json
// @Param object body RegisterBinder false "注册参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "注册成功"
// @failure 401 {object}  response.Information "邮箱已注册"
// @failure 402 {object}  response.Information "验证码错误"
// @failure 403 {object}  response.Information "请输入完整的信息"
// @Router /account/register [POST]
func UserRegister(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	code := c.PostForm("code")
	if len(email) == 0 || len(password) == 0 || len(code) == 0 {
		var reb RegisterBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.Json(c, 403, "请输入完整的信息", 0)
			return
		} else {
			email = reb.Email
			password = reb.Password
			code = reb.Code
		}
	}
	var ra account.Account
	user, err := ra.Register(email, password, code)
	if err == account.ExistUser {
		response.Json(c, 401, "邮箱已注册", 0)
		return
	} else if err == account.CodeErr {
		response.Json(c, 402, "验证码错误", 0)
		return
	} else {
		var mr redis.MailRedis
		mr.DelMail(email)
		response.Json(c, 200, "注册成功", user)
		return
	}
}

// GetMail 发送邮件
// @Summary 发送验证码邮件
// @Description 新用户发送验证码用于注册账号
// @Tags 验证码
// @Accept application/json
// @Produce application/json
// @Param object body MailBinder false "发送邮件参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "发送验证码成功"
// @failure 401 {object}  response.Information "邮件发送失败"
// @failure 402 {object}  response.Information "请输入邮箱或者密码"
// @Router /account/mail [POST]
func GetMail(c *gin.Context) {
	email := c.PostForm("email")
	if len(email) == 0 {
		var reb MailBinder
		if err := c.ShouldBind(&reb); err != nil {
			response.Json(c, 402, "请输入完整的信息", 0)
			return
		} else {
			email = reb.Email
		}
	}
	var gm account.Account
	code, err := gm.Mail(email)
	if err != nil {
		response.Json(c, 401, "邮件发送失败", err)
		return
	} else {
		var mr redis.MailRedis
		mr.SetMail(email, code)
		response.Json(c, 200, "邮件发送成功", code)
		return
	}
}

// UserLogout 用户注销
// @Summary 用户注销
// @Description 用于登录用户注销
// @Tags 注销
// @Accept application/json
// @Produce application/json
// @Param email query string false "注销参数"
// @Security ApiKeyAuth
// @Success 200 {object}  response.Information "注销成功"
// @failure 401 {object}  response.Information "注销失败"
// @failure 402 {object}  response.Information "您还未登录"
// @Router /account/logout [GET]
func UserLogout(c *gin.Context) {
	email := c.Query("email")
	if email == "" || len(email) == 0 {
		response.Json(c, 402, "请输入邮箱账号", 0)
		return
	}
	var la account.Account
	if err := la.Logout(email); err == nil {
		response.Json(c, 401, "注销失败", 0)
		return
	} else {
		var tr redis.TokenRedis
		tr.DelToken(email)
		response.Json(c, 200, "注销成功", 0)
		return
	}
}

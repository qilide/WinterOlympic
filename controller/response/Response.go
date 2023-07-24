package response

import "github.com/gin-gonic/gin"

type Response struct {
}

type Information struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type SuccessLogin struct {
	Information
	Token string `json:"token"`
}

func Json(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, Information{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func LoginJson(c *gin.Context, code int, msg string, data interface{}, token string) {
	c.JSON(code, SuccessLogin{
		Information: Information{
			Code: code,
			Msg:  msg,
			Data: data,
		},
		Token: token,
	})
}

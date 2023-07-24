package session

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context, email string) (string, error) {
	session := sessions.Default(c)
	if session.Get(email) != email {
		return "", errors.New("NotEmail")
	} else {
		return email, nil
	}
}

func DelCurrentUser(c *gin.Context, email string) {
	session := sessions.Default(c)
	session.Delete(email)
	if err := session.Save(); err != nil {
		fmt.Println("set user err...")
	}
}

func SetCurrentUser(c *gin.Context, email string) {
	session := sessions.Default(c)
	session.Set(email, email)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	if err := session.Save(); err != nil {
		fmt.Println("set user err...")
	}
}

func GetCurrentMail(c *gin.Context, email string) (string, error) {
	session := sessions.Default(c)
	code := session.Get(email + "mail")
	fmt.Println(code)
	if code != nil {
		fmt.Println("11111")
		return "", errors.New("NotCode")
	} else {
		fmt.Println("22222")
		return code.(string), nil
	}
}

func DelCurrentMail(c *gin.Context, email string) {
	session := sessions.Default(c)
	session.Delete(email)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	if err := session.Save(); err != nil {
		fmt.Println("del mail err...")
	}
}

func SetCurrentMail(c *gin.Context, email string, code string) {
	session := sessions.Default(c)
	session.Set(email+"mail", code)
	code1 := session.Get(email + "mail")
	fmt.Println(code1)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	if err := session.Save(); err != nil {
		fmt.Println("set mail err...")
	}
}

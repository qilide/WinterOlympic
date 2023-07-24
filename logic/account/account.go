package account

import (
	"WinterOlympic/dao/mysql"
	"WinterOlympic/dao/redis"
	"WinterOlympic/middlewares"
	"WinterOlympic/models"
	email2 "WinterOlympic/pkg/email"
	"WinterOlympic/pkg/snowflake"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Account struct {
}

var (
	NotRegister = errors.New("NotRegister") //该用户未注册
	PasswordErr = errors.New("PasswordErr") //邮箱或者密码错误
	ExistUser   = errors.New("ExistUser")   //邮箱已注册
	CodeErr     = errors.New("CodeErr")     //验证码错误
)

// 登录逻辑处理函数

func (la Account) Login(email string, password string) (models.User, error) {
	user := mysql.LoginVerity(email)
	if user.ID == 0 {
		return user, NotRegister
	} else {
		correct := middlewares.Check(password, user.Password)
		if correct == false {
			return user, PasswordErr
		} else {
			return user, nil
		}
	}
}

// 注册逻辑处理函数

func (ra Account) Register(email string, password string, code string) (models.User, error) {
	var user models.User
	var mr redis.MailRedis
	user = mysql.RegisterVerity(email)
	code1 := mr.GetMail(email)
	if user.ID != 0 {
		return user, ExistUser
	} else {
		if code1 != code {
			return user, CodeErr
		} else {
			var sf snowflake.Snowflake
			id := sf.NextVal()
			strInt64 := strconv.FormatInt(id, 10)
			id16, _ := strconv.Atoi(strInt64)
			currentTime := time.Now()
			pwd := middlewares.Encode(password)
			CurrentUser := mysql.CreateUser(id16, email, pwd, currentTime)
			return CurrentUser, nil
		}
	}
}

// 接收注册验证码逻辑处理函数

func (gm Account) Mail(email string) (string, error) {
	code, err := email2.SendMail(email)
	if err != nil {
		return "", err
	}
	return code, nil
}

// 注销逻辑处理函数

func (la Account) Logout(email string) interface{} {
	var tr redis.TokenRedis
	email1 := tr.GetToken(email)
	fmt.Println(email1)
	return email1
}

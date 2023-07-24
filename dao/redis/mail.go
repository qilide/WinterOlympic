package redis

import "time"

type MailRedis struct {
}

func (tr MailRedis) GetMail(email string) interface{} {
	token, _ := rdb.Do("Get", email+"mail").Result()
	return token
}

func (tr MailRedis) SetMail(email string, code string) {
	rdb.Set(email+"mail", code, time.Minute*10)
}

func (tr MailRedis) DelMail(email string) {
	rdb.Del(email + "mail")
}

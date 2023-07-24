package redis

import (
	"time"
)

type TokenRedis struct {
}

func (tr TokenRedis) GetToken(email string) interface{} {
	token, _ := rdb.Do("Get", email).Result()
	return token
}

func (tr TokenRedis) SetToken(email string, token string) {
	rdb.Set(email, token, time.Minute*30)
}

func (tr TokenRedis) DelToken(email string) {
	rdb.Del(email)
}

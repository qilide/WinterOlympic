package main

import (
	"WinterOlympic/dao/mysql"
	"WinterOlympic/dao/redis"
	"WinterOlympic/logger"
	"WinterOlympic/middlewares"
	"WinterOlympic/routes"
	"WinterOlympic/settings"
	"fmt"
	"go.uber.org/zap"
)

// @title 冬奥会智能分析与预测系统
// @version 1.0
// @description 本系统展示了前八界冬奥会数据和预测了下一届冬奥会各个国家获得的奖牌信息
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("Init settings failed, err: %v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("Init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化Mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("Init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化Redis连接
	if err := redis.Init(); err != nil {
		fmt.Printf("Init redis failed, err: %v\n", err)
		return
	}
	defer redis.Close()
	//5.注册路由
	r := routes.Setup()
	//6.启动服务(优雅关机)
	if err := middlewares.GracefulShutdown(r); err != nil {
		fmt.Printf("Graceful Shutdown failed, err: %v\n", err)
		return
	}
}

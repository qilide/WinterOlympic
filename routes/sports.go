package routes

import (
	"WinterOlympic/controller/sports"
	"github.com/gin-gonic/gin"
)

//注册主页面模块路由

func SportsGroupRoute(SportsGroup *gin.RouterGroup) {
	//国家奖牌数量排行
	SportsGroup.GET("/country_medal", sports.CountryMedal)
	//国家奖牌数量排行(详细)
	SportsGroup.GET("/country_medal_detail", sports.CountryMedalDetail)
	//项目奖牌数量排行
	SportsGroup.GET("/sports_medal", sports.SportsMedal)
	//项目奖牌数量排行(详细)
	SportsGroup.GET("/sports_medal_detail", sports.SportsMedalDetail)
	//国家信息
	SportsGroup.GET("/country_message", sports.CountryMessage)
	//运动员奖牌数量排行
	SportsGroup.GET("/athlete_medal", sports.AthleteMedal)
	//运动员奖牌数量排行(详细)
	SportsGroup.GET("/athlete_medal_detail", sports.AthleteMedalDetail)
	//中国项目奖牌数量排行
	SportsGroup.GET("/china_medal", sports.ChinaMedal)
	//中国项目奖牌数量排行(详细)
	SportsGroup.GET("/china_medal_detail", sports.ChinaMedalDetail)
}

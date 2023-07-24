package routes

import (
	"WinterOlympic/logger"
	"WinterOlympic/pkg/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func Setup() *gin.Engine {
	r := gin.New() //创立新的路由
	store := cookie.NewStore([]byte("WinterOlympic"))
	r.Static("/static", "./static")
	r.LoadHTMLFiles("templates/dist/index.html")
	r.Use(cors.Cors())
	r.Use(logger.GinLogger(), logger.GinRecovery(true), sessions.Sessions("WinterOlympic", store)) //使用日志记录路由信息
	pprof.Register(r)                                                                              //注册pprof相关路由
	InitSwagger(r)                                                                                 //注册Swagger文档路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	AccountGroupRoute(r.Group("/localApi/account/")) //注册账户功能模块路由
	SportsGroupRoute(r.Group("/localApi/sports/"))   //注册主页面模块路由
	CountryGroupRoute(r.Group("/localApi/country/")) //注册国家详细页面模块路由
	r.Run(":" + strconv.Itoa(viper.GetInt("port")))  //运行路由
	return r
}

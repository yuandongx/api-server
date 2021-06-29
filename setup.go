package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ping/modules/host"
	"ping/modules/setting"
	"regexp"
	"time"
)

func SetUp() *gin.Engine {
	g := gin.New()
	g.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "XMLHttpRequest", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			ko, _ := regexp.MatchString("https?://127.0.0.1:\\d+", origin)
			return ko
		},
		MaxAge: 12 * time.Hour,
	}))
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.ErrorMessage,
		)
	}))
	v1 := g.Group("/v1")
	{
		v1.POST("/cluster/:name/*action", host.V1Post)
		v1.GET("/cluster/:name/*id", host.V1Get)
		v1.GET("/setting/credential/*id", setting.AccessCredentGetV1)
	}
	return g
}

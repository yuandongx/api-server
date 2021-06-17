package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ping/modules/host"
	"regexp"
	"time"
)

func SetUp() *gin.Engine {
	g := gin.New()
	g.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS", "DELETE", "UPDATE"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "XMLHttpRequest", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			ko, _ := regexp.MatchString("https?://127.0.0.1:6060", origin)
			return ko
		},
		MaxAge: 12 * time.Hour,
	}))
	v1 := g.Group("/v1")
	{
		v1.POST("/cluster/:name/*action", host.V1Post)
		v1.GET("/cluster/:name/*id", host.V1Get)
	}
	return g
}

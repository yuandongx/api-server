package main

import (
	"github.com/gin-gonic/gin"
	"ping/modules/host"
)

func Test() {
	host.Host()
}

func SetUp() *gin.Engine  {
	g := gin.New()
	v1 := g.Group("/v1")
	{
		v1.POST("/host/:name/*action", host.Post)
		v1.POST("/host/:name/*id", host.Get)
	}
	return g
}
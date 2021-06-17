package host

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Host() {
	fmt.Println("11111111111111111")
}

func add(name string)  {

}

func del(name string)  {

}

func update(name string) {

}

func Post(c *gin.Context)  {
	name := c.Param("name")
	action := c.Param("action")
	if action == "add" {
		add(name)
	} else if action == "delete" {
		del(name)
	} else if action == "update" {
		update(name)
	}
}

func  Get(c *gin.Context)  {

}
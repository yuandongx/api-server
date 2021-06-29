package host

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ping/com/log"
	"ping/com/mysql"
	"strconv"
	"strings"
)

func add(name string) {

}

func del(name string) {

}

func update(name string) {

}

func V1Post(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	action = strings.Trim(action, "/")
	switch name {
	case "linux":
		obj := Linux{}
		c.ShouldBind(&obj)
		if action == "add" {
			code, err := mysql.Insert(&obj)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"msg": "failed", "code": "400"})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": code, "code": "200"})
			}
		} else if action == "delete" {
			del(name)
		} else if action == "update" {
			update(name)
		}
	case "router":
		obj := Router{}
		c.ShouldBind(&obj)
		if action == "add" {
			code, err := mysql.Insert(&obj)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"msg": "failed", "code": "400"})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": code, "code": "200"})
			}
		} else if action == "delete" {
			del(name)
		} else if action == "update" {
			update(name)
		}
	case "firewall":
		obj := Firewall{}
		c.ShouldBind(&obj)
		if action == "add" {
			code, err := mysql.Insert(&obj)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"msg": "failed", "code": "400"})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": code, "code": "200"})
			}
		} else if action == "delete" {
			del(name)
		} else if action == "update" {
			update(name)
		}
	case "switch":
		obj := Firewall{}
		c.ShouldBind(&obj)
		if action == "add" {
			code, err := mysql.Insert(&obj)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"msg": "failed", "code": "400"})
			} else {
				c.JSON(http.StatusOK, gin.H{"msg": code, "code": "200"})
			}
		} else if action == "delete" {
			del(name)
		} else if action == "update" {
			update(name)
		}
	}
}

func V1Get(c *gin.Context) {
	name := c.Param("name")
	id := c.Param("id")
	var obj mysql.Object
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"message": "no platform, no data!", "status": http.StatusOK})
	}
	id = strings.Trim(strings.Trim(id, "/"), " ")
	log.Display("--->", id, "<---")
	if id != "" {
		if Id, err := strconv.ParseInt(id, 10, 64); err == nil {
			switch name {
			case "switch":
				obj = &Switch{Id: Id}
			case "firewall":
				obj = &Firewall{}
			case "router":
				obj = &Router{}
			case "linux":
				obj = &Linux{}
			}
			mysql.Read(obj)
			c.JSON(http.StatusOK, obj)
		} else {
			log.Display("Error id: ", id)
			c.JSON(http.StatusOK, gin.H{"message": "error id", "status": http.StatusAccepted})
		}

	} else {
		switch name {
		case "switch":
			obj = &Switch{}
		case "firewall":
			obj = &Firewall{}
		case "router":
			obj = &Router{}
		case "linux":
			obj = &Linux{}
		}
		result := getAll(obj)
		c.JSON(http.StatusOK, result)
	}
}

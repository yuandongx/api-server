package setting

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AccessCredentGetV1(c *gin.Context) {
	id := c.Param("id")
	id = strings.Trim(strings.Trim(id, "/"), " ")
	if id != "" {
		a := read(id)
		access := a.(AccessCredentials)
		if access.Id == 0 {
			c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("No data found by the given id: %s", id), "status": 202})
		} else {
			c.JSON(http.StatusOK, access)
		}
	} else {
		result := fetchAll(AccessCredentials{})
		c.JSON(http.StatusOK, result)
	}
}

package apis

import (
	resource "github.com/nicelizhi/easy-admin/ui"

	"github.com/gin-gonic/gin"
)

func GoAdmin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, string(resource.Html))
}

func Favicon(c *gin.Context) {
	//c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, string(resource.Favicon))
}

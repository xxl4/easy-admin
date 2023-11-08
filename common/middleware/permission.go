package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2/util"
	ginI18n "github.com/gin-contrib/i18n"

	"github.com/nicelizhi/easy-admin-core/sdk"
	"github.com/nicelizhi/easy-admin-core/sdk/api"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	"github.com/gin-gonic/gin"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		e := sdk.Runtime.GetCasbinKey(c.Request.Host)
		var res, casbinExclude bool
		var err error
		if v["rolekey"] == "admin" {
			res = true
			c.Next()
			return
		}
		for _, i := range CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, 500, err, "")
			return
		}

		if res {
			log.Infof("isTrue: %v role: %s method: %s path: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, v["rolekey"], c.Request.Method, c.Request.URL.Path, ginI18n.MustGetMessage(c, "The current request does not have permission please check by the administrator"))
			c.JSON(http.StatusOK, gin.H{
				"code": 403,
				"msg":  ginI18n.MustGetMessage(c, "Sorry you do not have access to this interface, contact your administrator"),
			})
			c.Abort()
			return
		}

	}
}

package global

import (
	"github.com/nicelizhi/easy-admin-core/sdk"
	"github.com/nicelizhi/easy-admin-core/sdk/api"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func LoadPolicy(c *gin.Context) (*casbin.SyncedEnforcer, error) {
	log := api.GetRequestLogger(c)
	if err := sdk.Runtime.GetCasbinKey(c.Request.Host).LoadPolicy(); err == nil {
		return sdk.Runtime.GetCasbinKey(c.Request.Host), err
	} else {
		log.Errorf("casbin rbac_model or policy init error, %s ", err.Error())
		return nil, err
	}
}

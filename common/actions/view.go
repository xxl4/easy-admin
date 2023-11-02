package actions

import (
	"errors"
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"

	"github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	log "github.com/nicelizhi/easy-admin-core/logger"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nicelizhi/easy-admin/common/dto"
	"github.com/nicelizhi/easy-admin/common/models"
)

// ViewAction 通用详情动作
func ViewAction(control dto.Control, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		//查看详情
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, ginI18n.MustGetMessage(c, "Parameter validation failed"))
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Model generation failed"))
			return
		}

		var rsp interface{}
		if f != nil {
			rsp = f()
		} else {
			rsp, _ = req.GenerateM()
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.Model(object).WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Where(req.GetId()).First(rsp).Error

		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, http.StatusNotFound, nil, ginI18n.MustGetMessage(c, "The object being viewed does not exist or does not have permission to view it"))
			return
		}
		if err != nil {
			log.Errorf("MsgID[%s] View error: %s", msgID, err)
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "View failed"))
			return
		}
		response.OK(c, rsp, ginI18n.MustGetMessage(c, "Query successful"))
		c.Next()
	}
}

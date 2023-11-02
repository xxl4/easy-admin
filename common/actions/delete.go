package actions

import (
	"net/http"

	log "github.com/nicelizhi/easy-admin-core/logger"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth/user"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"

	"github.com/nicelizhi/easy-admin/common/dto"
	"github.com/nicelizhi/easy-admin/common/models"
)

// DeleteAction 通用删除动作
func DeleteAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		//删除操作
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
			response.Error(c, http.StatusUnprocessableEntity, err, ginI18n.MustGetMessage(c, "Parameter validation failed"))
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Model generation failed"))
			return
		}

		object.SetUpdateBy(user.GetUserId(c))

		//数据权限检查
		p := GetPermissionFromContext(c)

		db = db.WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Where(req.GetId()).Delete(object)
		if err = db.Error; err != nil {
			log.Errorf("MsgID[%s] Delete error: %s", msgID, err)
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Failed to delete"))
			return
		}
		if db.RowsAffected == 0 {
			response.Error(c, http.StatusForbidden, nil, ginI18n.MustGetMessage(c, "No right to delete this data"))
			return
		}
		response.OK(c, object.GetId(), ginI18n.MustGetMessage(c, "Successfully deleted"))
		c.Next()
	}
}

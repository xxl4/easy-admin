package actions

import (
	"errors"
	"net/http"

	log "github.com/nicelizhi/easy-admin-core/logger"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/nicelizhi/easy-admin/common/dto"
	"github.com/nicelizhi/easy-admin/common/models"
)

// IndexAction 通用查询动作
func IndexAction(m models.ActiveRecord, d dto.Index, f func() interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		list := f()
		object := m.Generate()
		req := d.Generate()
		var count int64

		//查询列表
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, ginI18n.MustGetMessage(c, "Parameter validation failed"))
			return
		}

		//数据权限检查
		p := GetPermissionFromContext(c)

		err = db.WithContext(c).Model(object).
			Scopes(
				dto.MakeCondition(req.GetNeedSearch()),
				dto.Paginate(req.GetPageSize(), req.GetPageIndex()),
				Permission(object.TableName(), p),
			).
			Find(list).Limit(-1).Offset(-1).
			Count(&count).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Errorf("MsgID[%s] Index error: %s", msgID, err)
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Query failed"))
			return
		}
		response.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), ginI18n.MustGetMessage(c, "Query successful"))
		c.Next()
	}
}

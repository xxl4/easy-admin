package actions

import (
	"net/http"

	"github.com/nicelizhi/easy-admin-core/sdk/api"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth/user"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"

	"github.com/nicelizhi/easy-admin/common/dto"
	"github.com/nicelizhi/easy-admin/common/models"
)

// CreateAction 通用新增动作
func CreateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		//新增操作
		req := control.Generate()
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, err.Error())
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Model generation failed"))
			return
		}
		object.SetCreateBy(user.GetUserId(c))
		err = db.WithContext(c).Create(object).Error
		if err != nil {
			log.Errorf("Create error: %s", err)
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Creation failed"))
			return
		}
		response.OK(c, object.GetId(), ginI18n.MustGetMessage(c, "Created successfully"))
		c.Next()
	}
}

package apis

import (
	"net/http"

	"github.com/nicelizhi/easy-admin-core/sdk"
	"github.com/nicelizhi/easy-admin-core/sdk/api"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"

	"github.com/nicelizhi/easy-admin/app/jobs/service"
	"github.com/nicelizhi/easy-admin/common/dto"
)

type SysJob struct {
	api.Api
}

// RemoveJobForService 调用service实现
func (e SysJob) RemoveJobForService(c *gin.Context) {
	v := dto.GeneralDelDto{}
	s := service.SysJob{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&v).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.RemoveJob(&v)
	if err != nil {
		e.Logger.Errorf("RemoveJob error, %s", err.Error())
		e.Error(500, err, "")
		return
	}
	e.OK(nil, s.Msg)
}

// StartJobForService 启动job service实现
func (e SysJob) StartJobForService(c *gin.Context) {
	e.MakeContext(c)
	log := e.GetLogger()
	db, err := e.GetOrm()
	if err != nil {
		log.Error(err)
		return
	}
	var v dto.GeneralGetDto
	err = c.BindUri(&v)
	if err != nil {
		log.Warnf("参数验证错误, error: %s", err)
		e.Error(http.StatusUnprocessableEntity, err, ginI18n.MustGetMessage(c, "Parameter validation failed"))
		return
	}
	s := service.SysJob{}
	s.Orm = db
	s.Log = log
	s.Cron = sdk.Runtime.GetCrontabKey(c.Request.Host)
	err = s.StartJob(&v)
	if err != nil {
		log.Errorf("GetCrontabKey error, %s", err.Error())
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, s.Msg)
}

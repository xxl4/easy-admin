package apis

import (
	"fmt"

	"github.com/nicelizhi/easy-admin/app/admin/models"
	"github.com/nicelizhi/easy-admin/app/admin/service"
	"github.com/nicelizhi/easy-admin/app/admin/service/dto"

	"github.com/nicelizhi/easy-admin-core/sdk/api"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SysOperaLog struct {
	api.Api
}

// GetPage 操作日志列表
// @Summary 操作日志列表
// @Description 获取JSON
// @Tags 操作日志
// @Param title query string false "title"
// @Param method query string false "method"
// @Param requestMethod  query string false "requestMethod"
// @Param operUrl query string false "operUrl"
// @Param operIp query string false "operIp"
// @Param status query string false "status"
// @Param beginTime query string false "beginTime"
// @Param endTime query string false "endTime"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log [get]
// @Security Bearer
func (e SysOperaLog) GetPage(c *gin.Context) {
	s := service.SysOperaLog{}
	req := new(dto.SysOperaLogGetPageReq)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	list := make([]models.SysOperaLog, 0)
	var count int64

	err = s.GetPage(req, &list, &count)
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), ginI18n.MustGetMessage(c, "Query successful"))
}

// Get 操作日志通过id获取
// @Summary 操作日志通过id获取
// @Description 获取JSON
// @Tags 操作日志
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log/{id} [get]
// @Security Bearer
func (e SysOperaLog) Get(c *gin.Context) {
	s := new(service.SysOperaLog)
	req := dto.SysOperaLogGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.SysOperaLog
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.OK(object, ginI18n.MustGetMessage(c, "Query successful"))
}

// Delete 操作日志删除
// DeleteSysMenu 操作日志删除
// @Summary 删除操作日志
// @Description 删除数据
// @Tags 操作日志
// @Param data body dto.SysOperaLogDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-opera-log [delete]
// @Security Bearer
func (e SysOperaLog) Delete(c *gin.Context) {
	s := new(service.SysOperaLog)
	req := dto.SysOperaLogDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(ginI18n.MustGetMessage(c, "Failed to delete! Error details"), err.Error()))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Successfully deleted"))
}

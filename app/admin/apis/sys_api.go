package apis

import (
	"github.com/nicelizhi/easy-admin-core/sdk/api"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/nicelizhi/easy-admin/app/admin/models"
	"github.com/nicelizhi/easy-admin/app/admin/service"
	"github.com/nicelizhi/easy-admin/app/admin/service/dto"
	"github.com/nicelizhi/easy-admin/common/actions"
)

type SysApi struct {
	api.Api
}

// GetPage 获取接口管理列表
// @Summary 获取接口管理列表
// @Description 获取接口管理列表
// @Tags 接口管理
// @Param name query string false "名称"
// @Param title query string false "标题"
// @Param path query string false "地址"
// @Param action query string false "类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SysApi}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-api [get]
// @Security Bearer
func (e SysApi) GetPage(c *gin.Context) {
	s := service.SysApi{}
	req := dto.SysApiGetPageReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.Form).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	//数据权限检查
	p := actions.GetPermissionFromContext(c)
	list := make([]models.SysApi, 0)
	var count int64
	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		//e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), ginI18n.MustGetMessage(c, "Query successful"))
}

// Get 获取接口管理
// @Summary 获取接口管理
// @Description 获取接口管理
// @Tags 接口管理
// @Param id path string false "id"
// @Success 200 {object} response.Response{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys-api/{id} [get]
// @Security Bearer
func (e SysApi) Get(c *gin.Context) {
	req := dto.SysApiGetReq{}
	s := service.SysApi{}
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
	p := actions.GetPermissionFromContext(c)
	var object models.SysApi
	err = s.Get(&req, p, &object).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.OK(object, ginI18n.MustGetMessage(c, "Query successful"))
}

// Update 修改接口管理
// @Summary 修改接口管理
// @Description 修改接口管理
// @Tags 接口管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": ginI18n.MustGetMessage(c, "Update completed")}"
// @Router /api/v1/sys-api/{id} [put]
// @Security Bearer
func (e SysApi) Update(c *gin.Context) {
	req := dto.SysApiUpdateReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)
	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Update failed"))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Update completed"))
}

// DeleteSysApi 删除接口管理
// @Summary 删除接口管理
// @Description 删除接口管理
// @Tags 接口管理
// @Param data body dto.SysApiDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": ginI18n.MustGetMessage(c, "Successfully deleted")}"
// @Router /api/v1/sys-api [delete]
// @Security Bearer
func (e SysApi) DeleteSysApi(c *gin.Context) {
	req := dto.SysApiDeleteReq{}
	s := service.SysApi{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	p := actions.GetPermissionFromContext(c)
	err = s.Remove(&req, p)
	if err != nil {
		//e.Error(500, err, ginI18n.MustGetMessage(c, "Failed to delete"))
		e.Error(500, err, ginI18n.MustGetMessage(c, "Failed to delete"))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Successfully deleted"))
}

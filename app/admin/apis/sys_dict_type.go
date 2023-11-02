package apis

import (
	"fmt"

	"github.com/nicelizhi/easy-admin/app/admin/models"

	"github.com/nicelizhi/easy-admin-core/sdk/api"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/nicelizhi/easy-admin-core/sdk/pkg/response"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/nicelizhi/easy-admin/app/admin/service"
	"github.com/nicelizhi/easy-admin/app/admin/service/dto"
)

type SysDictType struct {
	api.Api
}

// GetPage 字典类型列表数据
// @Summary 字典类型列表数据
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [get]
// @Security Bearer
func (e SysDictType) GetPage(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetPageReq{}
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
	list := make([]models.SysDictType, 0)
	var count int64
	err = s.GetPage(&req, &list, &count)
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), ginI18n.MustGetMessage(c, "Query successful"))
}

// Get 字典类型通过字典id获取
// @Summary 字典类型通过字典id获取
// @Description 获取JSON
// @Tags 字典类型
// @Param dictId path int true "字典类型编码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [get]
// @Security Bearer
func (e SysDictType) Get(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetReq{}
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
	var object models.SysDictType
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.OK(object, ginI18n.MustGetMessage(c, "Query successful"))
}

// Insert 字典类型创建
// @Summary 添加字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [post]
// @Security Bearer
func (e SysDictType) Insert(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeInsertReq{}
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
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(ginI18n.MustGetMessage(c, "Failed to create dictionary type details"), err.Error()))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Created successfully"))
}

// Update
// @Summary 修改字典类型
// @Description 获取JSON
// @Tags 字典类型
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysDictTypeUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type/{dictId} [put]
// @Security Bearer
func (e SysDictType) Update(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Error(500, err, err.Error())
		e.Logger.Error(err)
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Update completed"))
}

// Delete
// @Summary 删除字典类型
// @Description 删除数据
// @Tags 字典类型
// @Param dictCode body dto.SysDictTypeDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type [delete]
// @Security Bearer
func (e SysDictType) Delete(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	err = s.Remove(&req)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Successfully deleted"))
}

// GetAll
// @Summary 字典类型全部数据 代码生成使用接口
// @Description 获取JSON
// @Tags 字典类型
// @Param dictName query string false "dictName"
// @Param dictId query string false "dictId"
// @Param dictType query string false "dictType"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/type-option-select [get]
// @Security Bearer
func (e SysDictType) GetAll(c *gin.Context) {
	s := service.SysDictType{}
	req := dto.SysDictTypeGetPageReq{}
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
	list := make([]models.SysDictType, 0)
	err = s.GetAll(&req, &list)
	if err != nil {
		e.Error(500, err, err.Error())
		return
	}
	e.OK(list, ginI18n.MustGetMessage(c, "Query successful"))
}

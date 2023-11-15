package apis

import (
	"fmt"

	"github.com/nicelizhi/easy-admin/app/admin/models"

	"github.com/nicelizhi/easy-admin-core/sdk/api"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth/user"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	_ "github.com/nicelizhi/easy-admin-core/sdk/pkg/response"
	"github.com/nicelizhi/easy-admin/app/admin/service"
	"github.com/nicelizhi/easy-admin/app/admin/service/dto"
)

type SysMenu struct {
	api.Api
}

// GetPage Menu列表数据
// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [get]
// @Security Bearer
func (e SysMenu) GetPage(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuGetPageReq{}
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
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.OK(list, ginI18n.MustGetMessage(c, "Query successful"))
}

// GetPage Menu自定义列表列表数据
// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/cusmenu [get]
// @Security Bearer
func (e SysMenu) CusGetPage(c *gin.Context) {
	s := service.SysMenu{}
	req := dto.SysMenuGetPageReq{}
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
	var list = make([]models.SysMenu, 0)
	err = s.GetPage(&req, &list).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}

	list2 := make([]interface{}, 0)

	for index, item := range list {
		var litem = make(map[string]interface{})
		var meta = make(map[string]interface{})
		var children = make([]interface{}, 0)
		for index1, item2 := range item.Children {
			var litem2 = make(map[string]interface{})
			var meta2 = make(map[string]interface{})

			litem2["name"] = item2.MenuName
			litem2["path"] = item2.Path
			meta2["locale"] = item2.Permission
			meta2["requiresAuth"] = true
			meta2["hideInMenu"] = false
			meta2["icon"] = item2.Icon
			meta2["order"] = item2.Sort
			litem2["meta"] = meta2
			fmt.Println(index1)
			children = append(children[:len(children)], litem2)
		}

		meta["locale"] = item.Permission
		meta["requiresAuth"] = true
		meta["hideInMenu"] = false
		meta["icon"] = item.Icon
		meta["order"] = item.Sort
		litem["name"] = item.MenuName
		litem["path"] = item.Path
		litem["meta"] = meta
		litem["children"] = children
		fmt.Println(index, " = ", item)
		list2 = append(list2[:len(list2)], litem)
	}

	e.OK(list2, ginI18n.MustGetMessage(c, "Query successful"))
}

// Get 获取菜单详情
// @Summary Menu详情数据
// @Description 获取JSON
// @Tags 菜单
// @Param id path string false "id"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [get]
// @Security Bearer
func (e SysMenu) Get(c *gin.Context) {
	req := dto.SysMenuGetReq{}
	s := new(service.SysMenu)
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
	var object = models.SysMenu{}

	err = s.Get(&req, &object).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}
	e.OK(object, ginI18n.MustGetMessage(c, "Query successful"))
}

// Insert 创建菜单
// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param data body dto.SysMenuInsertReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [post]
// @Security Bearer
func (e SysMenu) Insert(c *gin.Context) {
	req := dto.SysMenuInsertReq{}
	s := new(service.SysMenu)
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Creation failed"))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Created successfully"))
}

// Update 修改菜单
// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SysMenuUpdateReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func (e SysMenu) Update(c *gin.Context) {
	req := dto.SysMenuUpdateReq{}
	s := new(service.SysMenu)
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
	err = s.Update(&req).Error
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Update failed"))
		return
	}
	e.OK(req.GetId(), ginI18n.MustGetMessage(c, "Update completed"))
}

// Delete 删除菜单
// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param data body dto.SysMenuDeleteReq true "body"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menu [delete]
// @Security Bearer
func (e SysMenu) Delete(c *gin.Context) {
	control := new(dto.SysMenuDeleteReq)
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(control, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	err = s.Remove(control).Error
	if err != nil {
		e.Logger.Errorf("RemoveSysMenu error, %s", err)
		e.Error(500, err, ginI18n.MustGetMessage(c, "Failed to delete"))
		return
	}
	e.OK(control.GetId(), ginI18n.MustGetMessage(c, "Successfully deleted"))
}

// GetMenuRole 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Summary 根据登录角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func (e SysMenu) GetMenuRole(c *gin.Context) {
	s := new(service.SysMenu)
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := s.SetMenuRole(user.GetRoleName(c))

	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}

	e.OK(result, "")
}

//// GetMenuIDS 获取角色对应的菜单id数组
//// @Summary 获取角色对应的菜单id数组，设置角色权限使用
//// @Description 获取JSON
//// @Tags 菜单
//// @Param id path int true "id"
//// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
//// @Router /api/v1/menuids/{id} [get]
//// @Security Bearer
//func (e SysMenu) GetMenuIDS(c *gin.Context) {
//	s := new(service.SysMenu)
//	r := service.SysRole{}
//	m := dto.SysRoleByName{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&m, binding.JSON).
//		MakeService(&s.Service).
//		MakeService(&r.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//	var data models.SysRole
//	err = r.GetWithName(&m, &data).Error
//
//	//data.RoleName = c.GetString("role")
//	//data.UpdateBy = user.GetUserId(c)
//	//result, err := data.GetIDS(s.Orm)
//
//	if err != nil {
//		e.Logger.Errorf("GetIDS error, %s", err.Error())
//		e.Error(500, err, "获取失败")
//		return
//	}
//	e.OK(result, "")
//}

// GetMenuTreeSelect 根据角色ID查询菜单下拉树结构
// @Summary 角色修改使用的菜单列表
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/json
// @Product application/json
// @Param roleId path int true "roleId"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/menuTreeselect/{roleId} [get]
// @Security Bearer
func (e SysMenu) GetMenuTreeSelect(c *gin.Context) {
	m := service.SysMenu{}
	r := service.SysRole{}
	req := dto.SelectRole{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&m.Service).
		MakeService(&r.Service).
		Bind(&req, nil).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	result, err := m.SetLabel()
	if err != nil {
		e.Error(500, err, ginI18n.MustGetMessage(c, "Query failed"))
		return
	}

	menuIds := make([]int, 0)
	if req.RoleId != 0 {
		menuIds, err = r.GetRoleMenuId(req.RoleId)
		if err != nil {
			e.Error(500, err, "")
			return
		}
	}
	e.OK(gin.H{
		"menus":       result,
		"checkedKeys": menuIds,
	}, "获取成功")
}

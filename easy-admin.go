package main

import (
	"github.com/nicelizhi/easy-admin/cmd"
)

// @title Easy-Admin
// @version 1.4.0
// @description 基于Gin + Vue 的前后端分离权限管理系统的接口文档
// @license.name Apache 2.0
// @license.url https://github.com/xxl4/easy-admin/blob/main/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}

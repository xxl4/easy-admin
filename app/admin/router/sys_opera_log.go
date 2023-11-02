package router

import (
	"github.com/nicelizhi/easy-admin/app/admin/apis"
	"github.com/nicelizhi/easy-admin/common/middleware"

	jwt "github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSysOperaLogRouter)
}

// 需认证的路由代码
func registerSysOperaLogRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysOperaLog{}
	r := v1.Group("/sys-opera-log").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.DELETE("", api.Delete)
	}
}

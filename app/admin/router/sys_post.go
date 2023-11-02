package router

import (
	"github.com/nicelizhi/easy-admin/app/admin/apis"
	"github.com/nicelizhi/easy-admin/common/middleware"

	jwt "github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerSyPostRouter)
}

// 需认证的路由代码
func registerSyPostRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysPost{}
	r := v1.Group("/post").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}

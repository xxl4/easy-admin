package router

import (
	"embed"
	"errors"
	"io/fs"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/nicelizhi/easy-admin/app/admin/apis"
	resource "github.com/nicelizhi/easy-admin/ui"

	"github.com/nicelizhi/easy-admin-core/sdk/config"

	jwt "github.com/nicelizhi/easy-admin-core/sdk/pkg/jwtauth"
	"github.com/nicelizhi/easy-admin-core/sdk/pkg/ws"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerfiles "github.com/swaggo/files"

	"github.com/nicelizhi/easy-admin/common/middleware"
	"github.com/nicelizhi/easy-admin/common/middleware/handler"
	_ "github.com/nicelizhi/easy-admin/docs/admin"
)

type Resource struct {
	fs   embed.FS
	path string
}

func CssResource() *Resource {
	return &Resource{
		fs:   resource.CssStatic,
		path: "dist/css",
	}
}

func JsResource() *Resource {
	return &Resource{
		fs:   resource.JsStatic,
		path: "dist/js",
	}
}

func ImgResource() *Resource {
	return &Resource{
		fs:   resource.ImgStatic,
		path: "dist/img",
	}
}

func FontsResource() *Resource {
	return &Resource{
		fs:   resource.FontStatic,
		path: "dist/fonts",
	}
}

func (r *Resource) Open(name string) (fs.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	fullName := filepath.Join(r.path, filepath.FromSlash(path.Clean("/"+name)))
	file, err := r.fs.Open(fullName)

	return file, err
}

func InitSysRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.RouterGroup {
	g := r.Group("")
	sysBaseRouter(g)
	// 静态文件
	sysStaticFileRouter(g)
	// swagger；注意：生产环境可以注释掉
	if config.ApplicationConfig.Mode != "prod" {
		sysSwaggerRouter(g)
	}
	// 需要认证
	sysCheckRoleRouterInit(g, authMiddleware)
	return g
}

func sysBaseRouter(r *gin.RouterGroup) {

	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendAllService()

	if config.ApplicationConfig.Mode != "prod" {
		r.GET("/", apis.GoAdmin)

		r.GET("/favicon.ico", apis.Favicon)
		r.StaticFS("/css", http.FS(CssResource()))
		r.StaticFS("/js", http.FS(JsResource()))
		r.StaticFS("/img", http.FS(ImgResource()))
		r.StaticFS("/fonts", http.FS(FontsResource()))

	}
	r.GET("/info", handler.Ping)
}

func sysStaticFileRouter(r *gin.RouterGroup) {
	err := mime.AddExtensionType(".js", "application/javascript")
	if err != nil {
		return
	}
	r.Static("/static", "./static")
	if config.ApplicationConfig.Mode != "prod" {
		r.Static("/form-generator", "./static/form-generator")
	}
}

func sysSwaggerRouter(r *gin.RouterGroup) {
	r.GET("/swagger/admin/*any", ginSwagger.WrapHandler(swaggerfiles.NewHandler(), ginSwagger.InstanceName("admin")))
}

func sysCheckRoleRouterInit(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	wss := r.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		wss.GET("/ws/:id/:channel", ws.WebsocketManager.WsClient)
		wss.GET("/wslogout/:id/:channel", ws.WebsocketManager.UnWsClient)
	}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)
		// Refresh time can be longer than token timeout
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	registerBaseRouter(v1, authMiddleware)
}

func registerBaseRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysMenu{}
	api2 := apis.SysDept{}
	v1auth := v1.Group("").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		v1auth.GET("/roleMenuTreeselect/:roleId", api.GetMenuTreeSelect)
		//v1.GET("/menuTreeselect", api.GetMenuTreeSelect)
		v1auth.GET("/roleDeptTreeselect/:roleId", api2.GetDeptTreeRoleSelect)
		v1auth.POST("/logout", handler.LogOut)
	}
}

package router

import (
	"os"

	common "github.com/nicelizhi/easy-admin/common/middleware"

	log "github.com/nicelizhi/easy-admin-core/logger"
	"github.com/nicelizhi/easy-admin-core/sdk"

	"github.com/gin-gonic/gin"
)

// InitRouter
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	// the jwt middleware
	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册系统路由
	InitSysRouter(r, authMiddleware)

	// 注册业务路由
	// TODO: 这里可存放业务路由，里边并无实际路由只有演示代码
	InitExamplesRouter(r, authMiddleware)
}

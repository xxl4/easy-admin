package server

import "github.com/nicelizhi/easy-admin/app/other/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}

package server

import "github.com/nicelizhi/easy-admin/app/jobs/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}

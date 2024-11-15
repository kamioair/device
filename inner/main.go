package main

import (
	"github.com/kamioair/qf/qservice"
)

func main() {
	// 创建微服务
	setting := qservice.NewSetting(DefModule, DefDesc, Version).
		BindInitFunc(onInit).
		BindReqFunc(onReqHandler).
		BindNoticeFunc(onNoticeHandler)
	service = qservice.NewService(setting)

	// 启动微服务
	service.Run()
}

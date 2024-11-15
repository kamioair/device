package main

import (
	"device/inner/blls"
	"device/inner/config"
	"device/inner/daos"
	"errors"
	"github.com/kamioair/qf/qdefine"
	"github.com/kamioair/qf/qservice"
	"github.com/kamioair/qf/utils/qconvert"
)

const (
	Version   = "V1.0.0"                     // 版本
	DefModule = "ClientManager"              // 模块名称
	DefDesc   = "客户端管理模块（用于客户端id分配、客户端信息维护）" // 模块描述
)

var (
	service *qservice.MicroService

	// 其他业务
	clientBll *blls.Client
)

// 初始化
func onInit(moduleName string) {
	// 配置初始化
	config.Init(moduleName)

	// 数据库初始化
	daos.Init(moduleName)

	// 业务初始化
	clientBll = blls.NewClient()
}

// 处理外部请求
func onReqHandler(route string, ctx qdefine.Context) (any, error) {
	switch route {
	case "NewDeviceCode":
		return clientBll.NewDeviceCode()
	case "GetDeviceList":
		return clientBll.GetDeviceList(ctx.GetString("key"))
	case "KnockDoor":
		return clientBll.KnockDoor(qconvert.ToAny[map[string]string](ctx.Raw()))
	}
	return nil, errors.New("route Not Matched")
}

// 处理外部通知
func onNoticeHandler(route string, ctx qdefine.Context) {

}

// 发送通知
func onNotice(route string, content any) {
	service.SendNotice(route, content)
}

// 发送日志
func onLog(logType qdefine.ELog, content string, err error) {
	service.SendLog(logType, content, err)
}

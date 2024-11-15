package daos

import "github.com/kamioair/qf/qdefine"

type ClientId struct {
	qdefine.DbSimple
	Type  string `gorm:"primaryKey"` // ID类型
	Value uint64 // ID值
}

type ClientInfo struct {
	qdefine.DbFull
	Name    string // 客户端名称
	Modules string // 包含的模块列表
}

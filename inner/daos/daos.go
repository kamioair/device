package daos

import (
	"github.com/kamioair/qf/qdefine"
	"github.com/kamioair/qf/utils/qdb"
)

var (
	IdDao   *qdefine.BaseDao[ClientId]
	InfoDao *qdefine.BaseDao[ClientInfo]
)

func Init(module string) {
	db := qdb.NewDb(module)

	// 初始化
	IdDao = qdefine.NewDao[ClientId](db)
	InfoDao = qdefine.NewDao[ClientInfo](db)
}

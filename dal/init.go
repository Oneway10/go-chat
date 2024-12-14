package dal

import (
	"chat/dal/dao"
	"chat/dal/mysql"
	"context"
)

func Init(ctx context.Context) {
	mysql.Init()
	dao.SetDefault(mysql.DB(context.Background()))
}

func GetTx(q []*dao.Query) *dao.Query {
	if len(q) == 1 && q[0] == nil {
		return dao.Q
	}
	if len(q) == 1 && q[0] != nil {
		return q[0]
	}
	return dao.Q
}

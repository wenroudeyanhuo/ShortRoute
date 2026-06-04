// @program:     ShotTener
// @file:        mysql.go
// @author:      16574
// @create:      2025-12-15 11:01
// @description:

package sequence

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

// 建立mysql连接 执行replace into 语句
// REPLACE INTO sequence (stub) VALUES ('a');
// SELECT LAST_INSERT_ID();
const sqlReplaceStub = `REPLACE INTO sequence (stub) VALUES ('a');`

type MySQL struct {
	conn sqlx.SqlConn
}

// 因为Mysql这个结构体实现了Next 方法也就是实现了Sequence 接口
func NewMySQL(dsn string) Sequence {
	return &MySQL{
		conn: sqlx.NewMysql(dsn),
	}
}

// Next 取下一个号
func (m *MySQL) Next() (seq uint64, err error) {
	//prepare 做准备
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceStub)
	if err != nil {
		logx.Errorw("conn.Prepare failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	defer stmt.Close()
	//执行
	execResult, err := stmt.Exec()
	if err != nil {
		logx.Errorw("stmt.Exec failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	//获取刚插入的主键id
	var lid int64
	lid, err = execResult.LastInsertId()
	if err != nil {
		logx.Errorw("execResult.LastInsertId failed", logx.LogField{Key: "err", Value: err.Error()})
		return 0, err
	}
	return uint64(lid), nil
}

package models

import (
	"database/sql"
)

// Account 用户信息结构
type Account struct {
	ID       int32          //账号id
	Name     string         //用户名
	Password string         //已加密的密码
	Email    sql.NullString //注册时填写的邮箱
	Point    int            //点数
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AccountUserBase is the golang structure of table account_user_base for DAO operations like Where/Data.
type AccountUserBase struct {
	g.Meta       `orm:"table:account_user_base, do:true"`
	UserId       interface{} // 用户编号
	UserAccount  interface{} // 用户账号
	UserPassword interface{} // 用户密码
	UserSalt     interface{} // salt值
}

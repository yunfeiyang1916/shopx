// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// AccountUserBase is the golang structure for table account_user_base.
type AccountUserBase struct {
	UserId       uint   `json:"userId"       orm:"user_id"       description:"用户编号"`  // 用户编号
	UserAccount  string `json:"userAccount"  orm:"user_account"  description:"用户账号"`  // 用户账号
	UserPassword string `json:"userPassword" orm:"user_password" description:"用户密码"`  // 用户密码
	UserSalt     string `json:"userSalt"     orm:"user_salt"     description:"salt值"` // salt值
}

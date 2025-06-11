// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AccountUserBaseDao is the data access object for the table account_user_base.
type AccountUserBaseDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  AccountUserBaseColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// AccountUserBaseColumns defines and stores column names for the table account_user_base.
type AccountUserBaseColumns struct {
	UserId       string // 用户编号
	UserAccount  string // 用户账号
	UserPassword string // 用户密码
	UserSalt     string // salt值
}

// accountUserBaseColumns holds the columns for the table account_user_base.
var accountUserBaseColumns = AccountUserBaseColumns{
	UserId:       "user_id",
	UserAccount:  "user_account",
	UserPassword: "user_password",
	UserSalt:     "user_salt",
}

// NewAccountUserBaseDao creates and returns a new DAO object for table data access.
func NewAccountUserBaseDao(handlers ...gdb.ModelHandler) *AccountUserBaseDao {
	return &AccountUserBaseDao{
		group:    "default",
		table:    "account_user_base",
		columns:  accountUserBaseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AccountUserBaseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AccountUserBaseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AccountUserBaseDao) Columns() AccountUserBaseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AccountUserBaseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AccountUserBaseDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AccountUserBaseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

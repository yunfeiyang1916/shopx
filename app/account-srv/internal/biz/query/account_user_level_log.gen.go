// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"shopx/app/account-srv/internal/biz/entity"
)

func newAccountUserLevelLog(db *gorm.DB, opts ...gen.DOOption) accountUserLevelLog {
	_accountUserLevelLog := accountUserLevelLog{}

	_accountUserLevelLog.accountUserLevelLogDo.UseDB(db, opts...)
	_accountUserLevelLog.accountUserLevelLogDo.UseModel(&entity.AccountUserLevelLog{})

	tableName := _accountUserLevelLog.accountUserLevelLogDo.TableName()
	_accountUserLevelLog.ALL = field.NewAsterisk(tableName)
	_accountUserLevelLog.UserLevelLogID = field.NewInt32(tableName, "user_level_log_id")
	_accountUserLevelLog.UserID = field.NewInt32(tableName, "user_id")
	_accountUserLevelLog.UserLevelID = field.NewInt32(tableName, "user_level_id")
	_accountUserLevelLog.UserLevelPreID = field.NewInt32(tableName, "user_level_pre_id")
	_accountUserLevelLog.OperateRemark = field.NewString(tableName, "operate_remark")
	_accountUserLevelLog.OperateUserID = field.NewInt32(tableName, "operate_user_id")
	_accountUserLevelLog.CreateTime = field.NewInt64(tableName, "create_time")

	_accountUserLevelLog.fillFieldMap()

	return _accountUserLevelLog
}

// accountUserLevelLog 用户等级记录表
type accountUserLevelLog struct {
	accountUserLevelLogDo accountUserLevelLogDo

	ALL            field.Asterisk
	UserLevelLogID field.Int32  // 日志编号
	UserID         field.Int32  // 用户编号
	UserLevelID    field.Int32  // 当前等级ID
	UserLevelPreID field.Int32  // 之前等级ID
	OperateRemark  field.String // 操作备注
	OperateUserID  field.Int32  // 操作人
	CreateTime     field.Int64  // 操作时间

	fieldMap map[string]field.Expr
}

func (a accountUserLevelLog) Table(newTableName string) *accountUserLevelLog {
	a.accountUserLevelLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountUserLevelLog) As(alias string) *accountUserLevelLog {
	a.accountUserLevelLogDo.DO = *(a.accountUserLevelLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountUserLevelLog) updateTableName(table string) *accountUserLevelLog {
	a.ALL = field.NewAsterisk(table)
	a.UserLevelLogID = field.NewInt32(table, "user_level_log_id")
	a.UserID = field.NewInt32(table, "user_id")
	a.UserLevelID = field.NewInt32(table, "user_level_id")
	a.UserLevelPreID = field.NewInt32(table, "user_level_pre_id")
	a.OperateRemark = field.NewString(table, "operate_remark")
	a.OperateUserID = field.NewInt32(table, "operate_user_id")
	a.CreateTime = field.NewInt64(table, "create_time")

	a.fillFieldMap()

	return a
}

func (a *accountUserLevelLog) WithContext(ctx context.Context) *accountUserLevelLogDo {
	return a.accountUserLevelLogDo.WithContext(ctx)
}

func (a accountUserLevelLog) TableName() string { return a.accountUserLevelLogDo.TableName() }

func (a accountUserLevelLog) Alias() string { return a.accountUserLevelLogDo.Alias() }

func (a accountUserLevelLog) Columns(cols ...field.Expr) gen.Columns {
	return a.accountUserLevelLogDo.Columns(cols...)
}

func (a *accountUserLevelLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountUserLevelLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["user_level_log_id"] = a.UserLevelLogID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["user_level_id"] = a.UserLevelID
	a.fieldMap["user_level_pre_id"] = a.UserLevelPreID
	a.fieldMap["operate_remark"] = a.OperateRemark
	a.fieldMap["operate_user_id"] = a.OperateUserID
	a.fieldMap["create_time"] = a.CreateTime
}

func (a accountUserLevelLog) clone(db *gorm.DB) accountUserLevelLog {
	a.accountUserLevelLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountUserLevelLog) replaceDB(db *gorm.DB) accountUserLevelLog {
	a.accountUserLevelLogDo.ReplaceDB(db)
	return a
}

type accountUserLevelLogDo struct{ gen.DO }

func (a accountUserLevelLogDo) Debug() *accountUserLevelLogDo {
	return a.withDO(a.DO.Debug())
}

func (a accountUserLevelLogDo) WithContext(ctx context.Context) *accountUserLevelLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountUserLevelLogDo) ReadDB() *accountUserLevelLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountUserLevelLogDo) WriteDB() *accountUserLevelLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountUserLevelLogDo) Session(config *gorm.Session) *accountUserLevelLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountUserLevelLogDo) Clauses(conds ...clause.Expression) *accountUserLevelLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountUserLevelLogDo) Returning(value interface{}, columns ...string) *accountUserLevelLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountUserLevelLogDo) Not(conds ...gen.Condition) *accountUserLevelLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountUserLevelLogDo) Or(conds ...gen.Condition) *accountUserLevelLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountUserLevelLogDo) Select(conds ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountUserLevelLogDo) Where(conds ...gen.Condition) *accountUserLevelLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountUserLevelLogDo) Order(conds ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountUserLevelLogDo) Distinct(cols ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountUserLevelLogDo) Omit(cols ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountUserLevelLogDo) Join(table schema.Tabler, on ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountUserLevelLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountUserLevelLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountUserLevelLogDo) Group(cols ...field.Expr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountUserLevelLogDo) Having(conds ...gen.Condition) *accountUserLevelLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountUserLevelLogDo) Limit(limit int) *accountUserLevelLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountUserLevelLogDo) Offset(offset int) *accountUserLevelLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountUserLevelLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountUserLevelLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountUserLevelLogDo) Unscoped() *accountUserLevelLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountUserLevelLogDo) Create(values ...*entity.AccountUserLevelLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountUserLevelLogDo) CreateInBatches(values []*entity.AccountUserLevelLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountUserLevelLogDo) Save(values ...*entity.AccountUserLevelLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountUserLevelLogDo) First() (*entity.AccountUserLevelLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserLevelLog), nil
	}
}

func (a accountUserLevelLogDo) Take() (*entity.AccountUserLevelLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserLevelLog), nil
	}
}

func (a accountUserLevelLogDo) Last() (*entity.AccountUserLevelLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserLevelLog), nil
	}
}

func (a accountUserLevelLogDo) Find() ([]*entity.AccountUserLevelLog, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AccountUserLevelLog), err
}

func (a accountUserLevelLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AccountUserLevelLog, err error) {
	buf := make([]*entity.AccountUserLevelLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountUserLevelLogDo) FindInBatches(result *[]*entity.AccountUserLevelLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountUserLevelLogDo) Attrs(attrs ...field.AssignExpr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountUserLevelLogDo) Assign(attrs ...field.AssignExpr) *accountUserLevelLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountUserLevelLogDo) Joins(fields ...field.RelationField) *accountUserLevelLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountUserLevelLogDo) Preload(fields ...field.RelationField) *accountUserLevelLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountUserLevelLogDo) FirstOrInit() (*entity.AccountUserLevelLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserLevelLog), nil
	}
}

func (a accountUserLevelLogDo) FirstOrCreate() (*entity.AccountUserLevelLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserLevelLog), nil
	}
}

func (a accountUserLevelLogDo) FindByPage(offset int, limit int) (result []*entity.AccountUserLevelLog, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a accountUserLevelLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountUserLevelLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountUserLevelLogDo) Delete(models ...*entity.AccountUserLevelLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountUserLevelLogDo) withDO(do gen.Dao) *accountUserLevelLogDo {
	a.DO = *do.(*gen.DO)
	return a
}

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

func newAccountUserHistory(db *gorm.DB, opts ...gen.DOOption) accountUserHistory {
	_accountUserHistory := accountUserHistory{}

	_accountUserHistory.accountUserHistoryDo.UseDB(db, opts...)
	_accountUserHistory.accountUserHistoryDo.UseModel(&entity.AccountUserHistory{})

	tableName := _accountUserHistory.accountUserHistoryDo.TableName()
	_accountUserHistory.ALL = field.NewAsterisk(tableName)
	_accountUserHistory.UserID = field.NewInt32(tableName, "user_id")
	_accountUserHistory.UserHistoryPassword = field.NewString(tableName, "user_history_password")
	_accountUserHistory.UserHistoryIP = field.NewString(tableName, "user_history_ip")
	_accountUserHistory.UserHistoryTime = field.NewInt64(tableName, "user_history_time")

	_accountUserHistory.fillFieldMap()

	return _accountUserHistory
}

// accountUserHistory 用户密码历史信息表
type accountUserHistory struct {
	accountUserHistoryDo accountUserHistoryDo

	ALL                 field.Asterisk
	UserID              field.Int32  // 用户编号
	UserHistoryPassword field.String // 历史密码
	UserHistoryIP       field.String // 历史 IP
	UserHistoryTime     field.Int64  // 修改时间

	fieldMap map[string]field.Expr
}

func (a accountUserHistory) Table(newTableName string) *accountUserHistory {
	a.accountUserHistoryDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountUserHistory) As(alias string) *accountUserHistory {
	a.accountUserHistoryDo.DO = *(a.accountUserHistoryDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountUserHistory) updateTableName(table string) *accountUserHistory {
	a.ALL = field.NewAsterisk(table)
	a.UserID = field.NewInt32(table, "user_id")
	a.UserHistoryPassword = field.NewString(table, "user_history_password")
	a.UserHistoryIP = field.NewString(table, "user_history_ip")
	a.UserHistoryTime = field.NewInt64(table, "user_history_time")

	a.fillFieldMap()

	return a
}

func (a *accountUserHistory) WithContext(ctx context.Context) *accountUserHistoryDo {
	return a.accountUserHistoryDo.WithContext(ctx)
}

func (a accountUserHistory) TableName() string { return a.accountUserHistoryDo.TableName() }

func (a accountUserHistory) Alias() string { return a.accountUserHistoryDo.Alias() }

func (a accountUserHistory) Columns(cols ...field.Expr) gen.Columns {
	return a.accountUserHistoryDo.Columns(cols...)
}

func (a *accountUserHistory) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountUserHistory) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 4)
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["user_history_password"] = a.UserHistoryPassword
	a.fieldMap["user_history_ip"] = a.UserHistoryIP
	a.fieldMap["user_history_time"] = a.UserHistoryTime
}

func (a accountUserHistory) clone(db *gorm.DB) accountUserHistory {
	a.accountUserHistoryDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountUserHistory) replaceDB(db *gorm.DB) accountUserHistory {
	a.accountUserHistoryDo.ReplaceDB(db)
	return a
}

type accountUserHistoryDo struct{ gen.DO }

func (a accountUserHistoryDo) Debug() *accountUserHistoryDo {
	return a.withDO(a.DO.Debug())
}

func (a accountUserHistoryDo) WithContext(ctx context.Context) *accountUserHistoryDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountUserHistoryDo) ReadDB() *accountUserHistoryDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountUserHistoryDo) WriteDB() *accountUserHistoryDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountUserHistoryDo) Session(config *gorm.Session) *accountUserHistoryDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountUserHistoryDo) Clauses(conds ...clause.Expression) *accountUserHistoryDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountUserHistoryDo) Returning(value interface{}, columns ...string) *accountUserHistoryDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountUserHistoryDo) Not(conds ...gen.Condition) *accountUserHistoryDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountUserHistoryDo) Or(conds ...gen.Condition) *accountUserHistoryDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountUserHistoryDo) Select(conds ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountUserHistoryDo) Where(conds ...gen.Condition) *accountUserHistoryDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountUserHistoryDo) Order(conds ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountUserHistoryDo) Distinct(cols ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountUserHistoryDo) Omit(cols ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountUserHistoryDo) Join(table schema.Tabler, on ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountUserHistoryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountUserHistoryDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountUserHistoryDo) Group(cols ...field.Expr) *accountUserHistoryDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountUserHistoryDo) Having(conds ...gen.Condition) *accountUserHistoryDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountUserHistoryDo) Limit(limit int) *accountUserHistoryDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountUserHistoryDo) Offset(offset int) *accountUserHistoryDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountUserHistoryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountUserHistoryDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountUserHistoryDo) Unscoped() *accountUserHistoryDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountUserHistoryDo) Create(values ...*entity.AccountUserHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountUserHistoryDo) CreateInBatches(values []*entity.AccountUserHistory, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountUserHistoryDo) Save(values ...*entity.AccountUserHistory) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountUserHistoryDo) First() (*entity.AccountUserHistory, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserHistory), nil
	}
}

func (a accountUserHistoryDo) Take() (*entity.AccountUserHistory, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserHistory), nil
	}
}

func (a accountUserHistoryDo) Last() (*entity.AccountUserHistory, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserHistory), nil
	}
}

func (a accountUserHistoryDo) Find() ([]*entity.AccountUserHistory, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AccountUserHistory), err
}

func (a accountUserHistoryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AccountUserHistory, err error) {
	buf := make([]*entity.AccountUserHistory, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountUserHistoryDo) FindInBatches(result *[]*entity.AccountUserHistory, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountUserHistoryDo) Attrs(attrs ...field.AssignExpr) *accountUserHistoryDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountUserHistoryDo) Assign(attrs ...field.AssignExpr) *accountUserHistoryDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountUserHistoryDo) Joins(fields ...field.RelationField) *accountUserHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountUserHistoryDo) Preload(fields ...field.RelationField) *accountUserHistoryDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountUserHistoryDo) FirstOrInit() (*entity.AccountUserHistory, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserHistory), nil
	}
}

func (a accountUserHistoryDo) FirstOrCreate() (*entity.AccountUserHistory, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserHistory), nil
	}
}

func (a accountUserHistoryDo) FindByPage(offset int, limit int) (result []*entity.AccountUserHistory, count int64, err error) {
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

func (a accountUserHistoryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountUserHistoryDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountUserHistoryDo) Delete(models ...*entity.AccountUserHistory) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountUserHistoryDo) withDO(do gen.Dao) *accountUserHistoryDo {
	a.DO = *do.(*gen.DO)
	return a
}

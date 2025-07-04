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

func newAccountUserBindConnect(db *gorm.DB, opts ...gen.DOOption) accountUserBindConnect {
	_accountUserBindConnect := accountUserBindConnect{}

	_accountUserBindConnect.accountUserBindConnectDo.UseDB(db, opts...)
	_accountUserBindConnect.accountUserBindConnectDo.UseModel(&entity.AccountUserBindConnect{})

	tableName := _accountUserBindConnect.accountUserBindConnectDo.TableName()
	_accountUserBindConnect.ALL = field.NewAsterisk(tableName)
	_accountUserBindConnect.BindID = field.NewString(tableName, "bind_id")
	_accountUserBindConnect.BindType = field.NewInt32(tableName, "bind_type")
	_accountUserBindConnect.UserID = field.NewInt32(tableName, "user_id")
	_accountUserBindConnect.BindTime = field.NewTime(tableName, "bind_time")
	_accountUserBindConnect.BindNickname = field.NewString(tableName, "bind_nickname")
	_accountUserBindConnect.BindIcon = field.NewString(tableName, "bind_icon")
	_accountUserBindConnect.BindOpenid = field.NewString(tableName, "bind_openid")
	_accountUserBindConnect.BindUnionid = field.NewString(tableName, "bind_unionid")
	_accountUserBindConnect.BindActive = field.NewField(tableName, "bind_active")

	_accountUserBindConnect.fillFieldMap()

	return _accountUserBindConnect
}

// accountUserBindConnect 用户绑定表
type accountUserBindConnect struct {
	accountUserBindConnectDo accountUserBindConnectDo

	ALL          field.Asterisk
	BindID       field.String // 绑定标记
	BindType     field.Int32  // 绑定类型(EMUN):1-mobile;  2-email;   13-weixin公众号
	UserID       field.Int32  // 用户编号
	BindTime     field.Time   // 绑定时间
	BindNickname field.String // 用户名称
	BindIcon     field.String // 用户图标
	BindOpenid   field.String // 访问编号
	BindUnionid  field.String // unionid
	BindActive   field.Field  // 是否激活(BOOL):0-未激活;1-激活

	fieldMap map[string]field.Expr
}

func (a accountUserBindConnect) Table(newTableName string) *accountUserBindConnect {
	a.accountUserBindConnectDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountUserBindConnect) As(alias string) *accountUserBindConnect {
	a.accountUserBindConnectDo.DO = *(a.accountUserBindConnectDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountUserBindConnect) updateTableName(table string) *accountUserBindConnect {
	a.ALL = field.NewAsterisk(table)
	a.BindID = field.NewString(table, "bind_id")
	a.BindType = field.NewInt32(table, "bind_type")
	a.UserID = field.NewInt32(table, "user_id")
	a.BindTime = field.NewTime(table, "bind_time")
	a.BindNickname = field.NewString(table, "bind_nickname")
	a.BindIcon = field.NewString(table, "bind_icon")
	a.BindOpenid = field.NewString(table, "bind_openid")
	a.BindUnionid = field.NewString(table, "bind_unionid")
	a.BindActive = field.NewField(table, "bind_active")

	a.fillFieldMap()

	return a
}

func (a *accountUserBindConnect) WithContext(ctx context.Context) *accountUserBindConnectDo {
	return a.accountUserBindConnectDo.WithContext(ctx)
}

func (a accountUserBindConnect) TableName() string { return a.accountUserBindConnectDo.TableName() }

func (a accountUserBindConnect) Alias() string { return a.accountUserBindConnectDo.Alias() }

func (a accountUserBindConnect) Columns(cols ...field.Expr) gen.Columns {
	return a.accountUserBindConnectDo.Columns(cols...)
}

func (a *accountUserBindConnect) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountUserBindConnect) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["bind_id"] = a.BindID
	a.fieldMap["bind_type"] = a.BindType
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["bind_time"] = a.BindTime
	a.fieldMap["bind_nickname"] = a.BindNickname
	a.fieldMap["bind_icon"] = a.BindIcon
	a.fieldMap["bind_openid"] = a.BindOpenid
	a.fieldMap["bind_unionid"] = a.BindUnionid
	a.fieldMap["bind_active"] = a.BindActive
}

func (a accountUserBindConnect) clone(db *gorm.DB) accountUserBindConnect {
	a.accountUserBindConnectDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountUserBindConnect) replaceDB(db *gorm.DB) accountUserBindConnect {
	a.accountUserBindConnectDo.ReplaceDB(db)
	return a
}

type accountUserBindConnectDo struct{ gen.DO }

func (a accountUserBindConnectDo) Debug() *accountUserBindConnectDo {
	return a.withDO(a.DO.Debug())
}

func (a accountUserBindConnectDo) WithContext(ctx context.Context) *accountUserBindConnectDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountUserBindConnectDo) ReadDB() *accountUserBindConnectDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountUserBindConnectDo) WriteDB() *accountUserBindConnectDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountUserBindConnectDo) Session(config *gorm.Session) *accountUserBindConnectDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountUserBindConnectDo) Clauses(conds ...clause.Expression) *accountUserBindConnectDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountUserBindConnectDo) Returning(value interface{}, columns ...string) *accountUserBindConnectDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountUserBindConnectDo) Not(conds ...gen.Condition) *accountUserBindConnectDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountUserBindConnectDo) Or(conds ...gen.Condition) *accountUserBindConnectDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountUserBindConnectDo) Select(conds ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountUserBindConnectDo) Where(conds ...gen.Condition) *accountUserBindConnectDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountUserBindConnectDo) Order(conds ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountUserBindConnectDo) Distinct(cols ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountUserBindConnectDo) Omit(cols ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountUserBindConnectDo) Join(table schema.Tabler, on ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountUserBindConnectDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountUserBindConnectDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountUserBindConnectDo) Group(cols ...field.Expr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountUserBindConnectDo) Having(conds ...gen.Condition) *accountUserBindConnectDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountUserBindConnectDo) Limit(limit int) *accountUserBindConnectDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountUserBindConnectDo) Offset(offset int) *accountUserBindConnectDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountUserBindConnectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountUserBindConnectDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountUserBindConnectDo) Unscoped() *accountUserBindConnectDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountUserBindConnectDo) Create(values ...*entity.AccountUserBindConnect) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountUserBindConnectDo) CreateInBatches(values []*entity.AccountUserBindConnect, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountUserBindConnectDo) Save(values ...*entity.AccountUserBindConnect) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountUserBindConnectDo) First() (*entity.AccountUserBindConnect, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserBindConnect), nil
	}
}

func (a accountUserBindConnectDo) Take() (*entity.AccountUserBindConnect, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserBindConnect), nil
	}
}

func (a accountUserBindConnectDo) Last() (*entity.AccountUserBindConnect, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserBindConnect), nil
	}
}

func (a accountUserBindConnectDo) Find() ([]*entity.AccountUserBindConnect, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AccountUserBindConnect), err
}

func (a accountUserBindConnectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AccountUserBindConnect, err error) {
	buf := make([]*entity.AccountUserBindConnect, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountUserBindConnectDo) FindInBatches(result *[]*entity.AccountUserBindConnect, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountUserBindConnectDo) Attrs(attrs ...field.AssignExpr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountUserBindConnectDo) Assign(attrs ...field.AssignExpr) *accountUserBindConnectDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountUserBindConnectDo) Joins(fields ...field.RelationField) *accountUserBindConnectDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountUserBindConnectDo) Preload(fields ...field.RelationField) *accountUserBindConnectDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountUserBindConnectDo) FirstOrInit() (*entity.AccountUserBindConnect, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserBindConnect), nil
	}
}

func (a accountUserBindConnectDo) FirstOrCreate() (*entity.AccountUserBindConnect, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserBindConnect), nil
	}
}

func (a accountUserBindConnectDo) FindByPage(offset int, limit int) (result []*entity.AccountUserBindConnect, count int64, err error) {
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

func (a accountUserBindConnectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountUserBindConnectDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountUserBindConnectDo) Delete(models ...*entity.AccountUserBindConnect) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountUserBindConnectDo) withDO(do gen.Dao) *accountUserBindConnectDo {
	a.DO = *do.(*gen.DO)
	return a
}

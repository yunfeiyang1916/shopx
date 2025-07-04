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

func newAccountUserMessage(db *gorm.DB, opts ...gen.DOOption) accountUserMessage {
	_accountUserMessage := accountUserMessage{}

	_accountUserMessage.accountUserMessageDo.UseDB(db, opts...)
	_accountUserMessage.accountUserMessageDo.UseModel(&entity.AccountUserMessage{})

	tableName := _accountUserMessage.accountUserMessageDo.TableName()
	_accountUserMessage.ALL = field.NewAsterisk(tableName)
	_accountUserMessage.MessageID = field.NewInt32(tableName, "message_id")
	_accountUserMessage.MessageParentID = field.NewInt32(tableName, "message_parent_id")
	_accountUserMessage.UserID = field.NewInt32(tableName, "user_id")
	_accountUserMessage.UserNickname = field.NewString(tableName, "user_nickname")
	_accountUserMessage.MessageKind = field.NewInt32(tableName, "message_kind")
	_accountUserMessage.UserOtherID = field.NewInt32(tableName, "user_other_id")
	_accountUserMessage.UserOtherNickname = field.NewString(tableName, "user_other_nickname")
	_accountUserMessage.MessageTitle = field.NewString(tableName, "message_title")
	_accountUserMessage.MessageContent = field.NewString(tableName, "message_content")
	_accountUserMessage.MessageTime = field.NewInt64(tableName, "message_time")
	_accountUserMessage.MessageIsRead = field.NewField(tableName, "message_is_read")
	_accountUserMessage.MessageIsDelete = field.NewField(tableName, "message_is_delete")
	_accountUserMessage.MessageType = field.NewInt32(tableName, "message_type")
	_accountUserMessage.MessageCat = field.NewString(tableName, "message_cat")
	_accountUserMessage.MessageDataType = field.NewInt32(tableName, "message_data_type")
	_accountUserMessage.MessageDataID = field.NewString(tableName, "message_data_id")
	_accountUserMessage.MessageLength = field.NewInt32(tableName, "message_length")
	_accountUserMessage.MessageW = field.NewInt32(tableName, "message_w")
	_accountUserMessage.MessageH = field.NewInt32(tableName, "message_h")

	_accountUserMessage.fillFieldMap()

	return _accountUserMessage
}

// accountUserMessage 短消息-聊天记录
type accountUserMessage struct {
	accountUserMessageDo accountUserMessageDo

	ALL               field.Asterisk
	MessageID         field.Int32  // 消息编号
	MessageParentID   field.Int32  // 上级编号
	UserID            field.Int32  // 所属用户:发送者或者接收者，如果message_kind=1则为当前用户发送的消息。
	UserNickname      field.String // 用户昵称
	MessageKind       field.Int32  // 消息种类(ENUM):1-发送消息;2-接收消息
	UserOtherID       field.Int32  // 相关用户:发送者或者接收者
	UserOtherNickname field.String // 相关昵称:发送者或者接收者
	MessageTitle      field.String // 消息标题
	MessageContent    field.String // 消息内容
	MessageTime       field.Int64  // 发送时间
	MessageIsRead     field.Field  // 是否读取(BOOL):0-未读;1-已读
	MessageIsDelete   field.Field  // 是否删除(BOOL):0-正常状态;1-删除状态
	MessageType       field.Int32  // 消息类型(ENUM):1-系统消息;2-用户消息
	MessageCat        field.String // 消息类型(ENUM):text-文本消息;img-图片消息;video-视频消息;file:文件;location:位置;redpack:红包
	MessageDataType   field.Int32  // 消息分类(ENUM):0-默认消息;1-公告消息;2-订单消息;3-商品消息;4-余额卡券;5-服务消息
	MessageDataID     field.String // 消息数据:商品编号|订单编号
	MessageLength     field.Int32  // 消息长度
	MessageW          field.Int32  // 图片宽度
	MessageH          field.Int32  // 图片高度

	fieldMap map[string]field.Expr
}

func (a accountUserMessage) Table(newTableName string) *accountUserMessage {
	a.accountUserMessageDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a accountUserMessage) As(alias string) *accountUserMessage {
	a.accountUserMessageDo.DO = *(a.accountUserMessageDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *accountUserMessage) updateTableName(table string) *accountUserMessage {
	a.ALL = field.NewAsterisk(table)
	a.MessageID = field.NewInt32(table, "message_id")
	a.MessageParentID = field.NewInt32(table, "message_parent_id")
	a.UserID = field.NewInt32(table, "user_id")
	a.UserNickname = field.NewString(table, "user_nickname")
	a.MessageKind = field.NewInt32(table, "message_kind")
	a.UserOtherID = field.NewInt32(table, "user_other_id")
	a.UserOtherNickname = field.NewString(table, "user_other_nickname")
	a.MessageTitle = field.NewString(table, "message_title")
	a.MessageContent = field.NewString(table, "message_content")
	a.MessageTime = field.NewInt64(table, "message_time")
	a.MessageIsRead = field.NewField(table, "message_is_read")
	a.MessageIsDelete = field.NewField(table, "message_is_delete")
	a.MessageType = field.NewInt32(table, "message_type")
	a.MessageCat = field.NewString(table, "message_cat")
	a.MessageDataType = field.NewInt32(table, "message_data_type")
	a.MessageDataID = field.NewString(table, "message_data_id")
	a.MessageLength = field.NewInt32(table, "message_length")
	a.MessageW = field.NewInt32(table, "message_w")
	a.MessageH = field.NewInt32(table, "message_h")

	a.fillFieldMap()

	return a
}

func (a *accountUserMessage) WithContext(ctx context.Context) *accountUserMessageDo {
	return a.accountUserMessageDo.WithContext(ctx)
}

func (a accountUserMessage) TableName() string { return a.accountUserMessageDo.TableName() }

func (a accountUserMessage) Alias() string { return a.accountUserMessageDo.Alias() }

func (a accountUserMessage) Columns(cols ...field.Expr) gen.Columns {
	return a.accountUserMessageDo.Columns(cols...)
}

func (a *accountUserMessage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *accountUserMessage) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 19)
	a.fieldMap["message_id"] = a.MessageID
	a.fieldMap["message_parent_id"] = a.MessageParentID
	a.fieldMap["user_id"] = a.UserID
	a.fieldMap["user_nickname"] = a.UserNickname
	a.fieldMap["message_kind"] = a.MessageKind
	a.fieldMap["user_other_id"] = a.UserOtherID
	a.fieldMap["user_other_nickname"] = a.UserOtherNickname
	a.fieldMap["message_title"] = a.MessageTitle
	a.fieldMap["message_content"] = a.MessageContent
	a.fieldMap["message_time"] = a.MessageTime
	a.fieldMap["message_is_read"] = a.MessageIsRead
	a.fieldMap["message_is_delete"] = a.MessageIsDelete
	a.fieldMap["message_type"] = a.MessageType
	a.fieldMap["message_cat"] = a.MessageCat
	a.fieldMap["message_data_type"] = a.MessageDataType
	a.fieldMap["message_data_id"] = a.MessageDataID
	a.fieldMap["message_length"] = a.MessageLength
	a.fieldMap["message_w"] = a.MessageW
	a.fieldMap["message_h"] = a.MessageH
}

func (a accountUserMessage) clone(db *gorm.DB) accountUserMessage {
	a.accountUserMessageDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a accountUserMessage) replaceDB(db *gorm.DB) accountUserMessage {
	a.accountUserMessageDo.ReplaceDB(db)
	return a
}

type accountUserMessageDo struct{ gen.DO }

func (a accountUserMessageDo) Debug() *accountUserMessageDo {
	return a.withDO(a.DO.Debug())
}

func (a accountUserMessageDo) WithContext(ctx context.Context) *accountUserMessageDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a accountUserMessageDo) ReadDB() *accountUserMessageDo {
	return a.Clauses(dbresolver.Read)
}

func (a accountUserMessageDo) WriteDB() *accountUserMessageDo {
	return a.Clauses(dbresolver.Write)
}

func (a accountUserMessageDo) Session(config *gorm.Session) *accountUserMessageDo {
	return a.withDO(a.DO.Session(config))
}

func (a accountUserMessageDo) Clauses(conds ...clause.Expression) *accountUserMessageDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a accountUserMessageDo) Returning(value interface{}, columns ...string) *accountUserMessageDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a accountUserMessageDo) Not(conds ...gen.Condition) *accountUserMessageDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a accountUserMessageDo) Or(conds ...gen.Condition) *accountUserMessageDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a accountUserMessageDo) Select(conds ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a accountUserMessageDo) Where(conds ...gen.Condition) *accountUserMessageDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a accountUserMessageDo) Order(conds ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a accountUserMessageDo) Distinct(cols ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a accountUserMessageDo) Omit(cols ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a accountUserMessageDo) Join(table schema.Tabler, on ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a accountUserMessageDo) LeftJoin(table schema.Tabler, on ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a accountUserMessageDo) RightJoin(table schema.Tabler, on ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a accountUserMessageDo) Group(cols ...field.Expr) *accountUserMessageDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a accountUserMessageDo) Having(conds ...gen.Condition) *accountUserMessageDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a accountUserMessageDo) Limit(limit int) *accountUserMessageDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a accountUserMessageDo) Offset(offset int) *accountUserMessageDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a accountUserMessageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *accountUserMessageDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a accountUserMessageDo) Unscoped() *accountUserMessageDo {
	return a.withDO(a.DO.Unscoped())
}

func (a accountUserMessageDo) Create(values ...*entity.AccountUserMessage) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a accountUserMessageDo) CreateInBatches(values []*entity.AccountUserMessage, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a accountUserMessageDo) Save(values ...*entity.AccountUserMessage) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a accountUserMessageDo) First() (*entity.AccountUserMessage, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserMessage), nil
	}
}

func (a accountUserMessageDo) Take() (*entity.AccountUserMessage, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserMessage), nil
	}
}

func (a accountUserMessageDo) Last() (*entity.AccountUserMessage, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserMessage), nil
	}
}

func (a accountUserMessageDo) Find() ([]*entity.AccountUserMessage, error) {
	result, err := a.DO.Find()
	return result.([]*entity.AccountUserMessage), err
}

func (a accountUserMessageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.AccountUserMessage, err error) {
	buf := make([]*entity.AccountUserMessage, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a accountUserMessageDo) FindInBatches(result *[]*entity.AccountUserMessage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a accountUserMessageDo) Attrs(attrs ...field.AssignExpr) *accountUserMessageDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a accountUserMessageDo) Assign(attrs ...field.AssignExpr) *accountUserMessageDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a accountUserMessageDo) Joins(fields ...field.RelationField) *accountUserMessageDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a accountUserMessageDo) Preload(fields ...field.RelationField) *accountUserMessageDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a accountUserMessageDo) FirstOrInit() (*entity.AccountUserMessage, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserMessage), nil
	}
}

func (a accountUserMessageDo) FirstOrCreate() (*entity.AccountUserMessage, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.AccountUserMessage), nil
	}
}

func (a accountUserMessageDo) FindByPage(offset int, limit int) (result []*entity.AccountUserMessage, count int64, err error) {
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

func (a accountUserMessageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a accountUserMessageDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a accountUserMessageDo) Delete(models ...*entity.AccountUserMessage) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *accountUserMessageDo) withDO(do gen.Dao) *accountUserMessageDo {
	a.DO = *do.(*gen.DO)
	return a
}

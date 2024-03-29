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

	"github.com/gogoclouds/gogo-services/gen/internal/admin/admin-service/internal/model"
)

func newAdmin(db *gorm.DB, opts ...gen.DOOption) admin {
	_admin := admin{}

	_admin.adminDo.UseDB(db, opts...)
	_admin.adminDo.UseModel(&model.Admin{})

	tableName := _admin.adminDo.TableName()
	_admin.ALL = field.NewAsterisk(tableName)
	_admin.ID = field.NewInt64(tableName, "id")
	_admin.Username = field.NewString(tableName, "username")
	_admin.Password = field.NewString(tableName, "password")
	_admin.Icon = field.NewString(tableName, "icon")
	_admin.Email = field.NewString(tableName, "email")
	_admin.Nickname = field.NewString(tableName, "nickname")
	_admin.Note = field.NewString(tableName, "note")
	_admin.LoginTime = field.NewTime(tableName, "login_time")
	_admin.Status = field.NewBool(tableName, "status")
	_admin.IsDel = field.NewField(tableName, "is_del")
	_admin.CreateTime = field.NewTime(tableName, "create_time")
	_admin.UpdateTime = field.NewTime(tableName, "update_time")

	_admin.fillFieldMap()

	return _admin
}

// admin 后台用户表
type admin struct {
	adminDo adminDo

	ALL        field.Asterisk
	ID         field.Int64
	Username   field.String
	Password   field.String
	Icon       field.String // 头像
	Email      field.String // 邮箱
	Nickname   field.String // 昵称
	Note       field.String // 备注信息
	LoginTime  field.Time   // 最后登录时间
	Status     field.Bool   // 帐号启用状态：0->禁用；1->启用
	IsDel      field.Field
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (a admin) Table(newTableName string) *admin {
	a.adminDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a admin) As(alias string) *admin {
	a.adminDo.DO = *(a.adminDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *admin) updateTableName(table string) *admin {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Username = field.NewString(table, "username")
	a.Password = field.NewString(table, "password")
	a.Icon = field.NewString(table, "icon")
	a.Email = field.NewString(table, "email")
	a.Nickname = field.NewString(table, "nickname")
	a.Note = field.NewString(table, "note")
	a.LoginTime = field.NewTime(table, "login_time")
	a.Status = field.NewBool(table, "status")
	a.IsDel = field.NewField(table, "is_del")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")

	a.fillFieldMap()

	return a
}

func (a *admin) WithContext(ctx context.Context) IAdminDo { return a.adminDo.WithContext(ctx) }

func (a admin) TableName() string { return a.adminDo.TableName() }

func (a admin) Alias() string { return a.adminDo.Alias() }

func (a admin) Columns(cols ...field.Expr) gen.Columns { return a.adminDo.Columns(cols...) }

func (a *admin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *admin) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["username"] = a.Username
	a.fieldMap["password"] = a.Password
	a.fieldMap["icon"] = a.Icon
	a.fieldMap["email"] = a.Email
	a.fieldMap["nickname"] = a.Nickname
	a.fieldMap["note"] = a.Note
	a.fieldMap["login_time"] = a.LoginTime
	a.fieldMap["status"] = a.Status
	a.fieldMap["is_del"] = a.IsDel
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
}

func (a admin) clone(db *gorm.DB) admin {
	a.adminDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a admin) replaceDB(db *gorm.DB) admin {
	a.adminDo.ReplaceDB(db)
	return a
}

type adminDo struct{ gen.DO }

type IAdminDo interface {
	gen.SubQuery
	Debug() IAdminDo
	WithContext(ctx context.Context) IAdminDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAdminDo
	WriteDB() IAdminDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAdminDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAdminDo
	Not(conds ...gen.Condition) IAdminDo
	Or(conds ...gen.Condition) IAdminDo
	Select(conds ...field.Expr) IAdminDo
	Where(conds ...gen.Condition) IAdminDo
	Order(conds ...field.Expr) IAdminDo
	Distinct(cols ...field.Expr) IAdminDo
	Omit(cols ...field.Expr) IAdminDo
	Join(table schema.Tabler, on ...field.Expr) IAdminDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo
	Group(cols ...field.Expr) IAdminDo
	Having(conds ...gen.Condition) IAdminDo
	Limit(limit int) IAdminDo
	Offset(offset int) IAdminDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo
	Unscoped() IAdminDo
	Create(values ...*model.Admin) error
	CreateInBatches(values []*model.Admin, batchSize int) error
	Save(values ...*model.Admin) error
	First() (*model.Admin, error)
	Take() (*model.Admin, error)
	Last() (*model.Admin, error)
	Find() ([]*model.Admin, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error)
	FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Admin) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAdminDo
	Assign(attrs ...field.AssignExpr) IAdminDo
	Joins(fields ...field.RelationField) IAdminDo
	Preload(fields ...field.RelationField) IAdminDo
	FirstOrInit() (*model.Admin, error)
	FirstOrCreate() (*model.Admin, error)
	FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAdminDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a adminDo) Debug() IAdminDo {
	return a.withDO(a.DO.Debug())
}

func (a adminDo) WithContext(ctx context.Context) IAdminDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a adminDo) ReadDB() IAdminDo {
	return a.Clauses(dbresolver.Read)
}

func (a adminDo) WriteDB() IAdminDo {
	return a.Clauses(dbresolver.Write)
}

func (a adminDo) Session(config *gorm.Session) IAdminDo {
	return a.withDO(a.DO.Session(config))
}

func (a adminDo) Clauses(conds ...clause.Expression) IAdminDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a adminDo) Returning(value interface{}, columns ...string) IAdminDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a adminDo) Not(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a adminDo) Or(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a adminDo) Select(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a adminDo) Where(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a adminDo) Order(conds ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a adminDo) Distinct(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a adminDo) Omit(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a adminDo) Join(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a adminDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a adminDo) RightJoin(table schema.Tabler, on ...field.Expr) IAdminDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a adminDo) Group(cols ...field.Expr) IAdminDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a adminDo) Having(conds ...gen.Condition) IAdminDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a adminDo) Limit(limit int) IAdminDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a adminDo) Offset(offset int) IAdminDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a adminDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAdminDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a adminDo) Unscoped() IAdminDo {
	return a.withDO(a.DO.Unscoped())
}

func (a adminDo) Create(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a adminDo) CreateInBatches(values []*model.Admin, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a adminDo) Save(values ...*model.Admin) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a adminDo) First() (*model.Admin, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Take() (*model.Admin, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Last() (*model.Admin, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) Find() ([]*model.Admin, error) {
	result, err := a.DO.Find()
	return result.([]*model.Admin), err
}

func (a adminDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Admin, err error) {
	buf := make([]*model.Admin, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a adminDo) FindInBatches(result *[]*model.Admin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a adminDo) Attrs(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a adminDo) Assign(attrs ...field.AssignExpr) IAdminDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a adminDo) Joins(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a adminDo) Preload(fields ...field.RelationField) IAdminDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a adminDo) FirstOrInit() (*model.Admin, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FirstOrCreate() (*model.Admin, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Admin), nil
	}
}

func (a adminDo) FindByPage(offset int, limit int) (result []*model.Admin, count int64, err error) {
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

func (a adminDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a adminDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a adminDo) Delete(models ...*model.Admin) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *adminDo) withDO(do gen.Dao) *adminDo {
	a.DO = *do.(*gen.DO)
	return a
}

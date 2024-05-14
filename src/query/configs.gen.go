// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
    "github.com/boostchicken/lol/model"
)

func newConfig(db *gorm.DB, opts ...gen.DOOption) config {
	_config := config{}

	_config.configDo.UseDB(db, opts...)
	_config.configDo.UseModel(&model.Config{})

	tableName := _config.configDo.TableName()
	_config.ALL = field.NewAsterisk(tableName)
	_config.Id = field.NewUint64(tableName, "id")
	_config.Tenant = field.NewString(tableName, "tenant")
	_config.Bind = field.NewString(tableName, "bind")
	_config.Entries = configHasManyEntries{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Entries", "model.LolEntry"),
	}

	_config.fillFieldMap()

	return _config
}

type config struct {
	configDo

	ALL     field.Asterisk
	Id      field.Uint64
	Tenant  field.String
	Bind    field.String
	Entries configHasManyEntries

	fieldMap map[string]field.Expr
}

func (c config) Table(newTableName string) *config {
	c.configDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c config) As(alias string) *config {
	c.configDo.DO = *(c.configDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *config) updateTableName(table string) *config {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewUint64(table, "id")
	c.Tenant = field.NewString(table, "tenant")
	c.Bind = field.NewString(table, "bind")

	c.fillFieldMap()

	return c
}

func (c *config) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *config) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 4)
	c.fieldMap["id"] = c.Id
	c.fieldMap["tenant"] = c.Tenant
	c.fieldMap["bind"] = c.Bind

}

func (c config) clone(db *gorm.DB) config {
	c.configDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c config) replaceDB(db *gorm.DB) config {
	c.configDo.ReplaceDB(db)
	return c
}

type configHasManyEntries struct {
	db *gorm.DB

	field.RelationField
}

func (a configHasManyEntries) Where(conds ...field.Expr) *configHasManyEntries {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a configHasManyEntries) WithContext(ctx context.Context) *configHasManyEntries {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a configHasManyEntries) Session(session *gorm.Session) *configHasManyEntries {
	a.db = a.db.Session(session)
	return &a
}

func (a configHasManyEntries) Model(m *model.Config) *configHasManyEntriesTx {
	return &configHasManyEntriesTx{a.db.Model(m).Association(a.Name())}
}

type configHasManyEntriesTx struct{ tx *gorm.Association }

func (a configHasManyEntriesTx) Find() (result []*model.LolEntry, err error) {
	return result, a.tx.Find(&result)
}

func (a configHasManyEntriesTx) Append(values ...*model.LolEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a configHasManyEntriesTx) Replace(values ...*model.LolEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a configHasManyEntriesTx) Delete(values ...*model.LolEntry) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a configHasManyEntriesTx) Clear() error {
	return a.tx.Clear()
}

func (a configHasManyEntriesTx) Count() int64 {
	return a.tx.Count()
}

type configDo struct{ gen.DO }

type IConfigDo interface {
	gen.SubQuery
	Debug() IConfigDo
	WithContext(ctx context.Context) IConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IConfigDo
	WriteDB() IConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IConfigDo
	Not(conds ...gen.Condition) IConfigDo
	Or(conds ...gen.Condition) IConfigDo
	Select(conds ...field.Expr) IConfigDo
	Where(conds ...gen.Condition) IConfigDo
	Order(conds ...field.Expr) IConfigDo
	Distinct(cols ...field.Expr) IConfigDo
	Omit(cols ...field.Expr) IConfigDo
	Join(table schema.Tabler, on ...field.Expr) IConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IConfigDo
	Group(cols ...field.Expr) IConfigDo
	Having(conds ...gen.Condition) IConfigDo
	Limit(limit int) IConfigDo
	Offset(offset int) IConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IConfigDo
	Unscoped() IConfigDo
	Create(values ...*model.Config) error
	CreateInBatches(values []*model.Config, batchSize int) error
	Save(values ...*model.Config) error
	First() (*model.Config, error)
	Take() (*model.Config, error)
	Last() (*model.Config, error)
	Find() ([]*model.Config, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Config, err error)
	FindInBatches(result *[]*model.Config, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Config) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IConfigDo
	Assign(attrs ...field.AssignExpr) IConfigDo
	Joins(fields ...field.RelationField) IConfigDo
	Preload(fields ...field.RelationField) IConfigDo
	FirstOrInit() (*model.Config, error)
	FirstOrCreate() (*model.Config, error)
	FindByPage(offset int, limit int) (result []*model.Config, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c configDo) Debug() IConfigDo {
	return c.withDO(c.DO.Debug())
}

func (c configDo) WithContext(ctx context.Context) IConfigDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c configDo) ReadDB() IConfigDo {
	return c.Clauses(dbresolver.Read)
}

func (c configDo) WriteDB() IConfigDo {
	return c.Clauses(dbresolver.Write)
}

func (c configDo) Session(config *gorm.Session) IConfigDo {
	return c.withDO(c.DO.Session(config))
}

func (c configDo) Clauses(conds ...clause.Expression) IConfigDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c configDo) Returning(value interface{}, columns ...string) IConfigDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c configDo) Not(conds ...gen.Condition) IConfigDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c configDo) Or(conds ...gen.Condition) IConfigDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c configDo) Select(conds ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c configDo) Where(conds ...gen.Condition) IConfigDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c configDo) Order(conds ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c configDo) Distinct(cols ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c configDo) Omit(cols ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c configDo) Join(table schema.Tabler, on ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c configDo) LeftJoin(table schema.Tabler, on ...field.Expr) IConfigDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c configDo) RightJoin(table schema.Tabler, on ...field.Expr) IConfigDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c configDo) Group(cols ...field.Expr) IConfigDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c configDo) Having(conds ...gen.Condition) IConfigDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c configDo) Limit(limit int) IConfigDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c configDo) Offset(offset int) IConfigDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c configDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IConfigDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c configDo) Unscoped() IConfigDo {
	return c.withDO(c.DO.Unscoped())
}

func (c configDo) Create(values ...*model.Config) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c configDo) CreateInBatches(values []*model.Config, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c configDo) Save(values ...*model.Config) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c configDo) First() (*model.Config, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Config), nil
	}
}

func (c configDo) Take() (*model.Config, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Config), nil
	}
}

func (c configDo) Last() (*model.Config, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Config), nil
	}
}

func (c configDo) Find() ([]*model.Config, error) {
	result, err := c.DO.Find()
	return result.([]*model.Config), err
}

func (c configDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Config, err error) {
	buf := make([]*model.Config, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c configDo) FindInBatches(result *[]*model.Config, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c configDo) Attrs(attrs ...field.AssignExpr) IConfigDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c configDo) Assign(attrs ...field.AssignExpr) IConfigDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c configDo) Joins(fields ...field.RelationField) IConfigDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c configDo) Preload(fields ...field.RelationField) IConfigDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c configDo) FirstOrInit() (*model.Config, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Config), nil
	}
}

func (c configDo) FirstOrCreate() (*model.Config, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Config), nil
	}
}

func (c configDo) FindByPage(offset int, limit int) (result []*model.Config, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c configDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c configDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c configDo) Delete(models ...*model.Config) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *configDo) withDO(do gen.Dao) *configDo {
	c.DO = *do.(*gen.DO)
	return c
}

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

	"ramblr/models"
)

func newBlog(db *gorm.DB, opts ...gen.DOOption) blog {
	_blog := blog{}

	_blog.blogDo.UseDB(db, opts...)
	_blog.blogDo.UseModel(&models.Blog{})

	tableName := _blog.blogDo.TableName()
	_blog.ALL = field.NewAsterisk(tableName)
	_blog.ID = field.NewUint(tableName, "id")
	_blog.CreatedAt = field.NewTime(tableName, "created_at")
	_blog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_blog.DeletedAt = field.NewField(tableName, "deleted_at")
	_blog.UserID = field.NewUint(tableName, "user_id")
	_blog.Url = field.NewString(tableName, "url")
	_blog.Title = field.NewString(tableName, "title")
	_blog.Description = field.NewString(tableName, "description")

	_blog.fillFieldMap()

	return _blog
}

type blog struct {
	blogDo

	ALL         field.Asterisk
	ID          field.Uint
	CreatedAt   field.Time
	UpdatedAt   field.Time
	DeletedAt   field.Field
	UserID      field.Uint
	Url         field.String
	Title       field.String
	Description field.String

	fieldMap map[string]field.Expr
}

func (b blog) Table(newTableName string) *blog {
	b.blogDo.UseTable(newTableName)
	return b.updateTableName(newTableName)
}

func (b blog) As(alias string) *blog {
	b.blogDo.DO = *(b.blogDo.As(alias).(*gen.DO))
	return b.updateTableName(alias)
}

func (b *blog) updateTableName(table string) *blog {
	b.ALL = field.NewAsterisk(table)
	b.ID = field.NewUint(table, "id")
	b.CreatedAt = field.NewTime(table, "created_at")
	b.UpdatedAt = field.NewTime(table, "updated_at")
	b.DeletedAt = field.NewField(table, "deleted_at")
	b.UserID = field.NewUint(table, "user_id")
	b.Url = field.NewString(table, "url")
	b.Title = field.NewString(table, "title")
	b.Description = field.NewString(table, "description")

	b.fillFieldMap()

	return b
}

func (b *blog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := b.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (b *blog) fillFieldMap() {
	b.fieldMap = make(map[string]field.Expr, 8)
	b.fieldMap["id"] = b.ID
	b.fieldMap["created_at"] = b.CreatedAt
	b.fieldMap["updated_at"] = b.UpdatedAt
	b.fieldMap["deleted_at"] = b.DeletedAt
	b.fieldMap["user_id"] = b.UserID
	b.fieldMap["url"] = b.Url
	b.fieldMap["title"] = b.Title
	b.fieldMap["description"] = b.Description
}

func (b blog) clone(db *gorm.DB) blog {
	b.blogDo.ReplaceConnPool(db.Statement.ConnPool)
	return b
}

func (b blog) replaceDB(db *gorm.DB) blog {
	b.blogDo.ReplaceDB(db)
	return b
}

type blogDo struct{ gen.DO }

type IBlogDo interface {
	gen.SubQuery
	Debug() IBlogDo
	WithContext(ctx context.Context) IBlogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IBlogDo
	WriteDB() IBlogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IBlogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IBlogDo
	Not(conds ...gen.Condition) IBlogDo
	Or(conds ...gen.Condition) IBlogDo
	Select(conds ...field.Expr) IBlogDo
	Where(conds ...gen.Condition) IBlogDo
	Order(conds ...field.Expr) IBlogDo
	Distinct(cols ...field.Expr) IBlogDo
	Omit(cols ...field.Expr) IBlogDo
	Join(table schema.Tabler, on ...field.Expr) IBlogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IBlogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IBlogDo
	Group(cols ...field.Expr) IBlogDo
	Having(conds ...gen.Condition) IBlogDo
	Limit(limit int) IBlogDo
	Offset(offset int) IBlogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IBlogDo
	Unscoped() IBlogDo
	Create(values ...*models.Blog) error
	CreateInBatches(values []*models.Blog, batchSize int) error
	Save(values ...*models.Blog) error
	First() (*models.Blog, error)
	Take() (*models.Blog, error)
	Last() (*models.Blog, error)
	Find() ([]*models.Blog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Blog, err error)
	FindInBatches(result *[]*models.Blog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Blog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IBlogDo
	Assign(attrs ...field.AssignExpr) IBlogDo
	Joins(fields ...field.RelationField) IBlogDo
	Preload(fields ...field.RelationField) IBlogDo
	FirstOrInit() (*models.Blog, error)
	FirstOrCreate() (*models.Blog, error)
	FindByPage(offset int, limit int) (result []*models.Blog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IBlogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (b blogDo) Debug() IBlogDo {
	return b.withDO(b.DO.Debug())
}

func (b blogDo) WithContext(ctx context.Context) IBlogDo {
	return b.withDO(b.DO.WithContext(ctx))
}

func (b blogDo) ReadDB() IBlogDo {
	return b.Clauses(dbresolver.Read)
}

func (b blogDo) WriteDB() IBlogDo {
	return b.Clauses(dbresolver.Write)
}

func (b blogDo) Session(config *gorm.Session) IBlogDo {
	return b.withDO(b.DO.Session(config))
}

func (b blogDo) Clauses(conds ...clause.Expression) IBlogDo {
	return b.withDO(b.DO.Clauses(conds...))
}

func (b blogDo) Returning(value interface{}, columns ...string) IBlogDo {
	return b.withDO(b.DO.Returning(value, columns...))
}

func (b blogDo) Not(conds ...gen.Condition) IBlogDo {
	return b.withDO(b.DO.Not(conds...))
}

func (b blogDo) Or(conds ...gen.Condition) IBlogDo {
	return b.withDO(b.DO.Or(conds...))
}

func (b blogDo) Select(conds ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Select(conds...))
}

func (b blogDo) Where(conds ...gen.Condition) IBlogDo {
	return b.withDO(b.DO.Where(conds...))
}

func (b blogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IBlogDo {
	return b.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (b blogDo) Order(conds ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Order(conds...))
}

func (b blogDo) Distinct(cols ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Distinct(cols...))
}

func (b blogDo) Omit(cols ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Omit(cols...))
}

func (b blogDo) Join(table schema.Tabler, on ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Join(table, on...))
}

func (b blogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IBlogDo {
	return b.withDO(b.DO.LeftJoin(table, on...))
}

func (b blogDo) RightJoin(table schema.Tabler, on ...field.Expr) IBlogDo {
	return b.withDO(b.DO.RightJoin(table, on...))
}

func (b blogDo) Group(cols ...field.Expr) IBlogDo {
	return b.withDO(b.DO.Group(cols...))
}

func (b blogDo) Having(conds ...gen.Condition) IBlogDo {
	return b.withDO(b.DO.Having(conds...))
}

func (b blogDo) Limit(limit int) IBlogDo {
	return b.withDO(b.DO.Limit(limit))
}

func (b blogDo) Offset(offset int) IBlogDo {
	return b.withDO(b.DO.Offset(offset))
}

func (b blogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IBlogDo {
	return b.withDO(b.DO.Scopes(funcs...))
}

func (b blogDo) Unscoped() IBlogDo {
	return b.withDO(b.DO.Unscoped())
}

func (b blogDo) Create(values ...*models.Blog) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Create(values)
}

func (b blogDo) CreateInBatches(values []*models.Blog, batchSize int) error {
	return b.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (b blogDo) Save(values ...*models.Blog) error {
	if len(values) == 0 {
		return nil
	}
	return b.DO.Save(values)
}

func (b blogDo) First() (*models.Blog, error) {
	if result, err := b.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blog), nil
	}
}

func (b blogDo) Take() (*models.Blog, error) {
	if result, err := b.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blog), nil
	}
}

func (b blogDo) Last() (*models.Blog, error) {
	if result, err := b.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blog), nil
	}
}

func (b blogDo) Find() ([]*models.Blog, error) {
	result, err := b.DO.Find()
	return result.([]*models.Blog), err
}

func (b blogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Blog, err error) {
	buf := make([]*models.Blog, 0, batchSize)
	err = b.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (b blogDo) FindInBatches(result *[]*models.Blog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return b.DO.FindInBatches(result, batchSize, fc)
}

func (b blogDo) Attrs(attrs ...field.AssignExpr) IBlogDo {
	return b.withDO(b.DO.Attrs(attrs...))
}

func (b blogDo) Assign(attrs ...field.AssignExpr) IBlogDo {
	return b.withDO(b.DO.Assign(attrs...))
}

func (b blogDo) Joins(fields ...field.RelationField) IBlogDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Joins(_f))
	}
	return &b
}

func (b blogDo) Preload(fields ...field.RelationField) IBlogDo {
	for _, _f := range fields {
		b = *b.withDO(b.DO.Preload(_f))
	}
	return &b
}

func (b blogDo) FirstOrInit() (*models.Blog, error) {
	if result, err := b.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blog), nil
	}
}

func (b blogDo) FirstOrCreate() (*models.Blog, error) {
	if result, err := b.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Blog), nil
	}
}

func (b blogDo) FindByPage(offset int, limit int) (result []*models.Blog, count int64, err error) {
	result, err = b.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = b.Offset(-1).Limit(-1).Count()
	return
}

func (b blogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = b.Count()
	if err != nil {
		return
	}

	err = b.Offset(offset).Limit(limit).Scan(result)
	return
}

func (b blogDo) Scan(result interface{}) (err error) {
	return b.DO.Scan(result)
}

func (b blogDo) Delete(models ...*models.Blog) (result gen.ResultInfo, err error) {
	return b.DO.Delete(models)
}

func (b *blogDo) withDO(do gen.Dao) *blogDo {
	b.DO = *do.(*gen.DO)
	return b
}

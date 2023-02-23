// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newSysUserRole(db *gorm.DB, opts ...gen.DOOption) sysUserRole {
	_sysUserRole := sysUserRole{}

	_sysUserRole.sysUserRoleDo.UseDB(db, opts...)
	_sysUserRole.sysUserRoleDo.UseModel(&models.SysUserRole{})

	tableName := _sysUserRole.sysUserRoleDo.TableName()
	_sysUserRole.ALL = field.NewAsterisk(tableName)
	_sysUserRole.ID = field.NewInt64(tableName, "id")
	_sysUserRole.UserID = field.NewInt64(tableName, "user_id")
	_sysUserRole.RoleID = field.NewInt64(tableName, "role_id")

	_sysUserRole.fillFieldMap()

	return _sysUserRole
}

type sysUserRole struct {
	sysUserRoleDo sysUserRoleDo

	ALL    field.Asterisk
	ID     field.Int64
	UserID field.Int64
	RoleID field.Int64

	fieldMap map[string]field.Expr
}

func (s sysUserRole) Table(newTableName string) *sysUserRole {
	s.sysUserRoleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUserRole) As(alias string) *sysUserRole {
	s.sysUserRoleDo.DO = *(s.sysUserRoleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUserRole) updateTableName(table string) *sysUserRole {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UserID = field.NewInt64(table, "user_id")
	s.RoleID = field.NewInt64(table, "role_id")

	s.fillFieldMap()

	return s
}

func (s *sysUserRole) WithContext(ctx context.Context) *sysUserRoleDo {
	return s.sysUserRoleDo.WithContext(ctx)
}

func (s sysUserRole) TableName() string { return s.sysUserRoleDo.TableName() }

func (s sysUserRole) Alias() string { return s.sysUserRoleDo.Alias() }

func (s *sysUserRole) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUserRole) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 3)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["role_id"] = s.RoleID
}

func (s sysUserRole) clone(db *gorm.DB) sysUserRole {
	s.sysUserRoleDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUserRole) replaceDB(db *gorm.DB) sysUserRole {
	s.sysUserRoleDo.ReplaceDB(db)
	return s
}

type sysUserRoleDo struct{ gen.DO }

func (s sysUserRoleDo) Debug() *sysUserRoleDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserRoleDo) WithContext(ctx context.Context) *sysUserRoleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserRoleDo) ReadDB() *sysUserRoleDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserRoleDo) WriteDB() *sysUserRoleDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserRoleDo) Session(config *gorm.Session) *sysUserRoleDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserRoleDo) Clauses(conds ...clause.Expression) *sysUserRoleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserRoleDo) Returning(value interface{}, columns ...string) *sysUserRoleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserRoleDo) Not(conds ...gen.Condition) *sysUserRoleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserRoleDo) Or(conds ...gen.Condition) *sysUserRoleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserRoleDo) Select(conds ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserRoleDo) Where(conds ...gen.Condition) *sysUserRoleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserRoleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysUserRoleDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysUserRoleDo) Order(conds ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserRoleDo) Distinct(cols ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserRoleDo) Omit(cols ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserRoleDo) Join(table schema.Tabler, on ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserRoleDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserRoleDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserRoleDo) Group(cols ...field.Expr) *sysUserRoleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserRoleDo) Having(conds ...gen.Condition) *sysUserRoleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserRoleDo) Limit(limit int) *sysUserRoleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserRoleDo) Offset(offset int) *sysUserRoleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserRoleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserRoleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserRoleDo) Unscoped() *sysUserRoleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserRoleDo) Create(values ...*models.SysUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserRoleDo) CreateInBatches(values []*models.SysUserRole, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserRoleDo) Save(values ...*models.SysUserRole) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserRoleDo) First() (*models.SysUserRole, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUserRole), nil
	}
}

func (s sysUserRoleDo) Take() (*models.SysUserRole, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUserRole), nil
	}
}

func (s sysUserRoleDo) Last() (*models.SysUserRole, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUserRole), nil
	}
}

func (s sysUserRoleDo) Find() ([]*models.SysUserRole, error) {
	result, err := s.DO.Find()
	return result.([]*models.SysUserRole), err
}

func (s sysUserRoleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysUserRole, err error) {
	buf := make([]*models.SysUserRole, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserRoleDo) FindInBatches(result *[]*models.SysUserRole, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserRoleDo) Attrs(attrs ...field.AssignExpr) *sysUserRoleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserRoleDo) Assign(attrs ...field.AssignExpr) *sysUserRoleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserRoleDo) Joins(fields ...field.RelationField) *sysUserRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserRoleDo) Preload(fields ...field.RelationField) *sysUserRoleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserRoleDo) FirstOrInit() (*models.SysUserRole, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUserRole), nil
	}
}

func (s sysUserRoleDo) FirstOrCreate() (*models.SysUserRole, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUserRole), nil
	}
}

func (s sysUserRoleDo) FindByPage(offset int, limit int) (result []*models.SysUserRole, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysUserRoleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserRoleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserRoleDo) Delete(models ...*models.SysUserRole) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserRoleDo) withDO(do gen.Dao) *sysUserRoleDo {
	s.DO = *do.(*gen.DO)
	return s
}

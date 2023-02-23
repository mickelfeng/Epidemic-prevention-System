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

func newSysDept(db *gorm.DB, opts ...gen.DOOption) sysDept {
	_sysDept := sysDept{}

	_sysDept.sysDeptDo.UseDB(db, opts...)
	_sysDept.sysDeptDo.UseModel(&models.SysDept{})

	tableName := _sysDept.sysDeptDo.TableName()
	_sysDept.ALL = field.NewAsterisk(tableName)
	_sysDept.DeptID = field.NewInt64(tableName, "dept_id")
	_sysDept.ParentID = field.NewInt64(tableName, "parent_id")
	_sysDept.DeptName = field.NewString(tableName, "dept_name")
	_sysDept.OrderNum = field.NewInt32(tableName, "order_num")
	_sysDept.Leader = field.NewString(tableName, "leader")
	_sysDept.Phone = field.NewString(tableName, "phone")
	_sysDept.Email = field.NewString(tableName, "email")
	_sysDept.Status = field.NewInt32(tableName, "status")
	_sysDept.CreateBy = field.NewString(tableName, "create_by")
	_sysDept.CreateTime = field.NewTime(tableName, "create_time")
	_sysDept.UpdateBy = field.NewString(tableName, "update_by")
	_sysDept.UpdateTime = field.NewTime(tableName, "update_time")
	_sysDept.IsDelete = field.NewInt32(tableName, "is_delete")
	_sysDept.Version = field.NewInt32(tableName, "version")

	_sysDept.fillFieldMap()

	return _sysDept
}

type sysDept struct {
	sysDeptDo sysDeptDo

	ALL        field.Asterisk
	DeptID     field.Int64  // 部门id
	ParentID   field.Int64  // 父部门id
	DeptName   field.String // 部门名称
	OrderNum   field.Int32  // 显示顺序
	Leader     field.String // 负责人
	Phone      field.String // 联系电话
	Email      field.String // 邮箱
	Status     field.Int32  // 部门状态（1正常 0停用）
	CreateBy   field.String // 创建者
	CreateTime field.Time   // 创建时间
	UpdateBy   field.String // 更新者
	UpdateTime field.Time   // 更新时间
	IsDelete   field.Int32
	Version    field.Int32

	fieldMap map[string]field.Expr
}

func (s sysDept) Table(newTableName string) *sysDept {
	s.sysDeptDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDept) As(alias string) *sysDept {
	s.sysDeptDo.DO = *(s.sysDeptDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDept) updateTableName(table string) *sysDept {
	s.ALL = field.NewAsterisk(table)
	s.DeptID = field.NewInt64(table, "dept_id")
	s.ParentID = field.NewInt64(table, "parent_id")
	s.DeptName = field.NewString(table, "dept_name")
	s.OrderNum = field.NewInt32(table, "order_num")
	s.Leader = field.NewString(table, "leader")
	s.Phone = field.NewString(table, "phone")
	s.Email = field.NewString(table, "email")
	s.Status = field.NewInt32(table, "status")
	s.CreateBy = field.NewString(table, "create_by")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateBy = field.NewString(table, "update_by")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.IsDelete = field.NewInt32(table, "is_delete")
	s.Version = field.NewInt32(table, "version")

	s.fillFieldMap()

	return s
}

func (s *sysDept) WithContext(ctx context.Context) *sysDeptDo { return s.sysDeptDo.WithContext(ctx) }

func (s sysDept) TableName() string { return s.sysDeptDo.TableName() }

func (s sysDept) Alias() string { return s.sysDeptDo.Alias() }

func (s *sysDept) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDept) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["dept_id"] = s.DeptID
	s.fieldMap["parent_id"] = s.ParentID
	s.fieldMap["dept_name"] = s.DeptName
	s.fieldMap["order_num"] = s.OrderNum
	s.fieldMap["leader"] = s.Leader
	s.fieldMap["phone"] = s.Phone
	s.fieldMap["email"] = s.Email
	s.fieldMap["status"] = s.Status
	s.fieldMap["create_by"] = s.CreateBy
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_by"] = s.UpdateBy
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["is_delete"] = s.IsDelete
	s.fieldMap["version"] = s.Version
}

func (s sysDept) clone(db *gorm.DB) sysDept {
	s.sysDeptDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDept) replaceDB(db *gorm.DB) sysDept {
	s.sysDeptDo.ReplaceDB(db)
	return s
}

type sysDeptDo struct{ gen.DO }

func (s sysDeptDo) Debug() *sysDeptDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDeptDo) WithContext(ctx context.Context) *sysDeptDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDeptDo) ReadDB() *sysDeptDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDeptDo) WriteDB() *sysDeptDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDeptDo) Session(config *gorm.Session) *sysDeptDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDeptDo) Clauses(conds ...clause.Expression) *sysDeptDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDeptDo) Returning(value interface{}, columns ...string) *sysDeptDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDeptDo) Not(conds ...gen.Condition) *sysDeptDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDeptDo) Or(conds ...gen.Condition) *sysDeptDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDeptDo) Select(conds ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDeptDo) Where(conds ...gen.Condition) *sysDeptDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDeptDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysDeptDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDeptDo) Order(conds ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDeptDo) Distinct(cols ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDeptDo) Omit(cols ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDeptDo) Join(table schema.Tabler, on ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDeptDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDeptDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDeptDo) Group(cols ...field.Expr) *sysDeptDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDeptDo) Having(conds ...gen.Condition) *sysDeptDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDeptDo) Limit(limit int) *sysDeptDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDeptDo) Offset(offset int) *sysDeptDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDeptDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysDeptDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDeptDo) Unscoped() *sysDeptDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDeptDo) Create(values ...*models.SysDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDeptDo) CreateInBatches(values []*models.SysDept, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDeptDo) Save(values ...*models.SysDept) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDeptDo) First() (*models.SysDept, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDept), nil
	}
}

func (s sysDeptDo) Take() (*models.SysDept, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDept), nil
	}
}

func (s sysDeptDo) Last() (*models.SysDept, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDept), nil
	}
}

func (s sysDeptDo) Find() ([]*models.SysDept, error) {
	result, err := s.DO.Find()
	return result.([]*models.SysDept), err
}

func (s sysDeptDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysDept, err error) {
	buf := make([]*models.SysDept, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDeptDo) FindInBatches(result *[]*models.SysDept, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDeptDo) Attrs(attrs ...field.AssignExpr) *sysDeptDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDeptDo) Assign(attrs ...field.AssignExpr) *sysDeptDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDeptDo) Joins(fields ...field.RelationField) *sysDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDeptDo) Preload(fields ...field.RelationField) *sysDeptDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDeptDo) FirstOrInit() (*models.SysDept, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDept), nil
	}
}

func (s sysDeptDo) FirstOrCreate() (*models.SysDept, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysDept), nil
	}
}

func (s sysDeptDo) FindByPage(offset int, limit int) (result []*models.SysDept, count int64, err error) {
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

func (s sysDeptDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDeptDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDeptDo) Delete(models ...*models.SysDept) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDeptDo) withDO(do gen.Dao) *sysDeptDo {
	s.DO = *do.(*gen.DO)
	return s
}

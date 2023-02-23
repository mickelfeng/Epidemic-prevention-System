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

func newSysUser(db *gorm.DB, opts ...gen.DOOption) sysUser {
	_sysUser := sysUser{}

	_sysUser.sysUserDo.UseDB(db, opts...)
	_sysUser.sysUserDo.UseModel(&models.SysUser{})

	tableName := _sysUser.sysUserDo.TableName()
	_sysUser.ALL = field.NewAsterisk(tableName)
	_sysUser.ID = field.NewInt64(tableName, "id")
	_sysUser.Username = field.NewString(tableName, "username")
	_sysUser.Password = field.NewString(tableName, "password")
	_sysUser.Nickname = field.NewString(tableName, "nickname")
	_sysUser.Avatar = field.NewString(tableName, "avatar")
	_sysUser.PhoneNumber = field.NewString(tableName, "phone_number")
	_sysUser.City = field.NewString(tableName, "city")
	_sysUser.DeptID = field.NewInt64(tableName, "dept_id")
	_sysUser.CreateTime = field.NewTime(tableName, "create_time")
	_sysUser.UpdateTime = field.NewTime(tableName, "update_time")
	_sysUser.Remark = field.NewString(tableName, "remark")
	_sysUser.Status = field.NewInt32(tableName, "status")
	_sysUser.IsDelete = field.NewInt32(tableName, "is_delete")
	_sysUser.Version = field.NewInt32(tableName, "version")

	_sysUser.fillFieldMap()

	return _sysUser
}

type sysUser struct {
	sysUserDo sysUserDo

	ALL         field.Asterisk
	ID          field.Int64
	Username    field.String
	Password    field.String
	Nickname    field.String // 昵称
	Avatar      field.String
	PhoneNumber field.String
	City        field.String
	DeptID      field.Int64
	CreateTime  field.Time
	UpdateTime  field.Time
	Remark      field.String
	Status      field.Int32
	IsDelete    field.Int32
	Version     field.Int32

	fieldMap map[string]field.Expr
}

func (s sysUser) Table(newTableName string) *sysUser {
	s.sysUserDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysUser) As(alias string) *sysUser {
	s.sysUserDo.DO = *(s.sysUserDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysUser) updateTableName(table string) *sysUser {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.Username = field.NewString(table, "username")
	s.Password = field.NewString(table, "password")
	s.Nickname = field.NewString(table, "nickname")
	s.Avatar = field.NewString(table, "avatar")
	s.PhoneNumber = field.NewString(table, "phone_number")
	s.City = field.NewString(table, "city")
	s.DeptID = field.NewInt64(table, "dept_id")
	s.CreateTime = field.NewTime(table, "create_time")
	s.UpdateTime = field.NewTime(table, "update_time")
	s.Remark = field.NewString(table, "remark")
	s.Status = field.NewInt32(table, "status")
	s.IsDelete = field.NewInt32(table, "is_delete")
	s.Version = field.NewInt32(table, "version")

	s.fillFieldMap()

	return s
}

func (s *sysUser) WithContext(ctx context.Context) *sysUserDo { return s.sysUserDo.WithContext(ctx) }

func (s sysUser) TableName() string { return s.sysUserDo.TableName() }

func (s sysUser) Alias() string { return s.sysUserDo.Alias() }

func (s *sysUser) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysUser) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 14)
	s.fieldMap["id"] = s.ID
	s.fieldMap["username"] = s.Username
	s.fieldMap["password"] = s.Password
	s.fieldMap["nickname"] = s.Nickname
	s.fieldMap["avatar"] = s.Avatar
	s.fieldMap["phone_number"] = s.PhoneNumber
	s.fieldMap["city"] = s.City
	s.fieldMap["dept_id"] = s.DeptID
	s.fieldMap["create_time"] = s.CreateTime
	s.fieldMap["update_time"] = s.UpdateTime
	s.fieldMap["remark"] = s.Remark
	s.fieldMap["status"] = s.Status
	s.fieldMap["is_delete"] = s.IsDelete
	s.fieldMap["version"] = s.Version
}

func (s sysUser) clone(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysUser) replaceDB(db *gorm.DB) sysUser {
	s.sysUserDo.ReplaceDB(db)
	return s
}

type sysUserDo struct{ gen.DO }

func (s sysUserDo) Debug() *sysUserDo {
	return s.withDO(s.DO.Debug())
}

func (s sysUserDo) WithContext(ctx context.Context) *sysUserDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysUserDo) ReadDB() *sysUserDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysUserDo) WriteDB() *sysUserDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysUserDo) Session(config *gorm.Session) *sysUserDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysUserDo) Clauses(conds ...clause.Expression) *sysUserDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysUserDo) Returning(value interface{}, columns ...string) *sysUserDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysUserDo) Not(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysUserDo) Or(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysUserDo) Select(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysUserDo) Where(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysUserDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysUserDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysUserDo) Order(conds ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysUserDo) Distinct(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysUserDo) Omit(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysUserDo) Join(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysUserDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysUserDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysUserDo) Group(cols ...field.Expr) *sysUserDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysUserDo) Having(conds ...gen.Condition) *sysUserDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysUserDo) Limit(limit int) *sysUserDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysUserDo) Offset(offset int) *sysUserDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysUserDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysUserDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysUserDo) Unscoped() *sysUserDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysUserDo) Create(values ...*models.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysUserDo) CreateInBatches(values []*models.SysUser, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysUserDo) Save(values ...*models.SysUser) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysUserDo) First() (*models.SysUser, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUser), nil
	}
}

func (s sysUserDo) Take() (*models.SysUser, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUser), nil
	}
}

func (s sysUserDo) Last() (*models.SysUser, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUser), nil
	}
}

func (s sysUserDo) Find() ([]*models.SysUser, error) {
	result, err := s.DO.Find()
	return result.([]*models.SysUser), err
}

func (s sysUserDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.SysUser, err error) {
	buf := make([]*models.SysUser, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysUserDo) FindInBatches(result *[]*models.SysUser, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysUserDo) Attrs(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysUserDo) Assign(attrs ...field.AssignExpr) *sysUserDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysUserDo) Joins(fields ...field.RelationField) *sysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysUserDo) Preload(fields ...field.RelationField) *sysUserDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysUserDo) FirstOrInit() (*models.SysUser, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUser), nil
	}
}

func (s sysUserDo) FirstOrCreate() (*models.SysUser, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.SysUser), nil
	}
}

func (s sysUserDo) FindByPage(offset int, limit int) (result []*models.SysUser, count int64, err error) {
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

func (s sysUserDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysUserDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysUserDo) Delete(models ...*models.SysUser) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysUserDo) withDO(do gen.Dao) *sysUserDo {
	s.DO = *do.(*gen.DO)
	return s
}

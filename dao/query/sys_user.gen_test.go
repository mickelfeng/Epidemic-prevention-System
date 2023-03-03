// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"github.com/pwh-pwh/Epidemic-prevention-System/models"
	"testing"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&models.SysUser{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&models.SysUser{}) fail: %s", err)
	}
}

func Test_ZoreUp(t *testing.T) {
	userQ := Use(db).SysUser
	u := new(models.SysUser)
	userQ.WithContext(context.Background()).Create(u)
}

func Test_at(t *testing.T) {
	userQ := Use(db).SysUser
	u := new(models.SysUser)
	u.ID = 6
	userQ.WithContext(context.Background()).Debug().Select(userQ.ALL).Where(userQ.ID.Eq(6)).Updates(u)
}

func Test_u(t *testing.T) {
	u := new(models.SysUser)
	u.ID = 6
	db.Debug().Model(u).Select("*").Omit("create_time").Where("id=?", 6).Updates(u)
}

func Test_sysUserQuery(t *testing.T) {
	sysUser := newSysUser(db)
	sysUser = *sysUser.As(sysUser.TableName())
	_do := sysUser.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(sysUser.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <sys_user> fail:", err)
		return
	}

	_, ok := sysUser.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from sysUser success")
	}

	err = _do.Create(&models.SysUser{})
	if err != nil {
		t.Error("create item in table <sys_user> fail:", err)
	}

	err = _do.Save(&models.SysUser{})
	if err != nil {
		t.Error("create item in table <sys_user> fail:", err)
	}

	err = _do.CreateInBatches([]*models.SysUser{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <sys_user> fail:", err)
	}

	_, err = _do.Select(sysUser.ALL).Take()
	if err != nil {
		t.Error("Take() on table <sys_user> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <sys_user> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <sys_user> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <sys_user> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*models.SysUser{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <sys_user> fail:", err)
	}

	_, err = _do.Select(sysUser.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <sys_user> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <sys_user> fail:", err)
	}

	_, err = _do.Select(sysUser.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <sys_user> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <sys_user> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <sys_user> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <sys_user> fail:", err)
	}

	_, err = _do.ScanByPage(&models.SysUser{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <sys_user> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <sys_user> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <sys_user> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <sys_user> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <sys_user> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <sys_user> fail:", err)
	}
}

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameSysUserRole = "sys_user_role"

// SysUserRole mapped from table <sys_user_role>
type SysUserRole struct {
	ID     int64 `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	UserID int64 `gorm:"column:user_id;type:bigint(20);not null" json:"userId"`
	RoleID int64 `gorm:"column:role_id;type:bigint(20);not null" json:"roleId"`
}

// TableName SysUserRole's table name
func (*SysUserRole) TableName() string {
	return TableNameSysUserRole
}

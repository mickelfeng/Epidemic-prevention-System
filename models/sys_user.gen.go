// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"encoding"
	"encoding/json"
	"gorm.io/plugin/soft_delete"
)

const TableNameSysUser = "sys_user"

// SysUser mapped from table <sys_user>
type SysUser struct {
	ID          int64                 `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"`
	Username    string                `gorm:"column:username;type:varchar(64);uniqueIndex:UK_USERNAME,priority:1" json:"username"`
	Password    string                `gorm:"column:password;type:varchar(64)" json:"password"`
	Nickname    string                `gorm:"column:nickname;type:varchar(64)" json:"nickname"` // 昵称
	Avatar      string                `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	PhoneNumber string                `gorm:"column:phone_number;type:varchar(64)" json:"phoneNumber"`
	City        string                `gorm:"column:city;type:varchar(64)" json:"city"`
	DeptID      int64                 `gorm:"column:dept_id;type:bigint(20)" json:"dept_id"`
	CreateTime  LocalTime             `gorm:"column:create_time;type:datetime;autoCreateTime:true" json:"createTime"` // 创建时间
	UpdateTime  LocalTime             `gorm:"column:update_time;type:datetime;autoUpdateTime:true" json:"updateTime"` // 更新时间
	IsDelete    soft_delete.DeletedAt `gorm:"column:is_delete;type:int(1);softDelete:flag" json:"-"`                  // 逻辑删除
	Remark      string                `gorm:"column:remark;type:varchar(50)" json:"remark"`
	Status      int32                 `gorm:"column:status;type:int(5);not null" json:"status"`
	Version     int32                 `gorm:"column:version;type:int(11);default:1" json:"version"`
}

func (u *SysUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *SysUser) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

var _ encoding.BinaryMarshaler = new(SysUser)
var _ encoding.BinaryUnmarshaler = new(SysUser)

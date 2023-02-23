// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"time"
)

const TableNameHealthClock = "health_clock"

// HealthClock mapped from table <health_clock>
type HealthClock struct {
	ID          int64     `gorm:"column:id;type:bigint(11);primaryKey;autoIncrement:true" json:"id"` // 打卡id
	Username    string    `gorm:"column:username;type:varchar(50)" json:"username"`                  // 姓名
	HealthType  int32     `gorm:"column:health_type;type:int(1)" json:"health_type"`                 // 健康状况
	Temperature float32   `gorm:"column:temperature;type:float(6,1)" json:"temperature"`             // 温度
	MiddleHigh  int32     `gorm:"column:middle_high;type:int(1)" json:"middle_high"`                 // 中高风险
	Diagnosis   int32     `gorm:"column:diagnosis;type:int(1)" json:"diagnosis"`                     // 确诊
	ReturnInfo  int32     `gorm:"column:return_info;type:int(1)" json:"return_info"`                 // 境外返回
	Address     string    `gorm:"column:address;type:varchar(60)" json:"address"`                    // 地址
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`               // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`               // 更新时间
	DeptID      int32     `gorm:"column:dept_id;type:int(11)" json:"dept_id"`                        // 部门id
	IsDelete    int32     `gorm:"column:is_delete;type:int(1)" json:"is_delete"`                     // 逻辑删除
}

// TableName HealthClock's table name
func (*HealthClock) TableName() string {
	return TableNameHealthClock
}
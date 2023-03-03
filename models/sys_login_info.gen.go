// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

const TableNameSysLoginInfo = "sys_login_info"

// SysLoginInfo mapped from table <sys_login_info>
type SysLoginInfo struct {
	ID            int64     `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true" json:"id"` // 访问ID
	Username      string    `gorm:"column:username;type:varchar(50)" json:"username"`                  // 用户账号
	IP            string    `gorm:"column:ip;type:varchar(50)" json:"ip"`                              // 登录IP地址
	LoginLocation string    `gorm:"column:login_location;type:varchar(100)" json:"loginLocation"`      // 登录地点
	Browser       string    `gorm:"column:browser;type:varchar(50)" json:"browser"`                    // 浏览器类型
	Os            string    `gorm:"column:os;type:varchar(50)" json:"os"`                              // 操作系统
	Status        int32     `gorm:"column:status;type:int(1);default:1" json:"status"`                 // 登录状态（0成功 1失败）
	Msg           string    `gorm:"column:msg;type:varchar(100)" json:"msg"`                           // 提示消息
	LoginTime     LocalTime `gorm:"column:login_time;type:datetime" json:"loginTime"`                  // 访问时间
}

// TableName SysLoginInfo's table name
func (*SysLoginInfo) TableName() string {
	return TableNameSysLoginInfo
}

package table

import "time"

// UserPermissionType bit
type UserPermissionType int

// User the user table
type User struct {
	UID         int                `xorm:"user_uid not null pk autoincr INT(20)"`
	UIN         string             `xorm:"id not null comment('用户ID名') unique VARCHAR(64)"`
	Pwd         string             `xorm:"pwd not null comment('密码') VARCHAR(128)"`
	Salt        string             `xorm:"salt not null comment('盐值') VARCHAR(128)"`
	Role        string             `xorm:"role comment('角色') VARCHAR(128)"`
	Desc        string             `xorm:"desc comment('描述') VARCHAR(128)"`
	Permisssion UserPermissionType `xorm:"account_type not null comment('账户类型 123...') TINYINT(4)"`

	CreatedAt time.Time `xorm:"ctime not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP created"`
	UpdatedAt time.Time `xorm:"utime not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP updated"`
	Version   int       `xorm:"version"`
}

// TableName return the table name
func (t *User) TableName() string {
	return "t_user"
}

// Set Get Is Can

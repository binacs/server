package table

import "time"

// Page the page table
type Page struct {
	ID         int           `xorm:"pid not null pk autoincr INT(20)"`
	Poster     string        `xorm:"poster not null comment('Poster') VARCHAR(64)"`
	Syntax     string        `xorm:"syntax not null comment('Syntax') VARCHAR(64)"`
	Content    string        `xorm:"content not null comment('Content') TEXT(65535)"`
	TinyURL    string        `xorm:"tinyurl not null comment('TinyURL') VARCHAR(64)"`
	Expiration time.Duration `xorm:"expir not null comment('Expiration') BIGINT(30)"`

	CreatedAt time.Time `xorm:"ctime not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP created"`
	UpdatedAt time.Time `xorm:"utime not null default CURRENT_TIMESTAMP comment('更新时间') TIMESTAMP updated"`
	Version   int       `xorm:"version"`
}

// TableName return the table name
func (t *Page) TableName() string {
	return "t_page"
}

// Set Get Is Can

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "user"

// User 用户表
type User struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name       string    `gorm:"column:name;not null;comment:用户名" json:"name"`                                          // 用户名
	Password   string    `gorm:"column:password;not null;comment:密码" json:"password"`                                   // 密码
	Email      *string   `gorm:"column:email;comment:邮箱" json:"email"`                                                  // 邮箱
	Phone      *string   `gorm:"column:phone;comment:手机号" json:"phone"`                                                 // 手机号
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"` // 更新时间
	IsDeleted  int32     `gorm:"column:is_deleted;not null;comment:0-未删除,1-已删除" json:"is_deleted"`                      // 0-未删除,1-已删除
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

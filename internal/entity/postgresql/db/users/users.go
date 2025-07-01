package users

import (
	"ai_writer/internal/interactor/models/special"
)

// Table struct is users database table struct
type Table struct {
	// 使用者名稱
	UserName string `gorm:"column:user_name;type:text;" json:"user_name"`
	// 使用者中文名稱
	Name string `gorm:"column:name;type:text;" json:"name"`
	// 使用者密碼
	Password string `gorm:"column:password;type:text;" json:"password"`
	// 使用者電話
	PhoneNumber string `gorm:"column:phone_number;type:text;" json:"phone_number"`
	// 是否啟用
	Active bool `gorm:"column:active;type:boolean;default:true;" json:"active"`
	// 點數
	Point int `gorm:"column:point;type:integer;" json:"point"`
	// 使用者電子郵件
	Email string `gorm:"column:email;type:text;" json:"email"`
	special.Table
}

// Base struct is corresponding to users table structure file
type Base struct {
	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者密碼
	Password *string `json:"password,omitempty"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 是否啟用
	Active *bool `json:"active,omitempty"`
	// 點數
	Point *int `json:"point,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"`
	special.Base
}

// TableName sets the insert table name for this struct type
func (t *Table) TableName() string {
	return "users"
}

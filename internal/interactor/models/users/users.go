package users

import (
	"ai_writer/internal/interactor/models/page"
	"ai_writer/internal/interactor/models/section"
)

// Create struct is used to create achieves
type Create struct {
	// 使用者名稱
	UserName string `json:"user_name,omitempty" binding:"required" validate:"required"`
	// 使用者中文名稱
	Name string `json:"name,omitempty" binding:"required" validate:"required"`
	// 使用者密碼
	Password string `json:"password,omitempty" binding:"required" validate:"required"`
	// 使用者電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 是否啟用
	Active bool `json:"active,omitempty" default:"true"`
	// 點數
	Point int `json:"point,omitempty" default:"0"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
	// 角色ID
	RoleID string `json:"role_id,omitempty" `
	// 創建者
	CreatedBy string `json:"created_by,omitempty" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`

	// 使用者名稱
	UserName *string `json:"user_name,omitempty" form:"user_name"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty" form:"name"`
	// 使用者密碼
	Password *string `json:"password,omitempty" form:"password"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty" form:"phone_number"`
	// 是否啟用
	Active *bool `json:"active,omitempty" form:"active"`
	// 點數
	Point *int `json:"point,omitempty" form:"point"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty" form:"email"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	// 搜尋結構檔
	Field
	// 分頁搜尋結構檔
	page.Pagination
}

// List is multiple return structure files
type List struct {
	// 多筆
	Users []*struct {
		// 表ID
		ID string `json:"id,omitempty"`

		// 使用者名稱
		UserName string `json:"user_name,omitempty"`
		// 使用者中文名稱
		Name string `json:"name,omitempty"`
		// 使用者電話
		PhoneNumber string `json:"phone_number,omitempty"`
		// 是否啟用
		Active bool `json:"active,omitempty"`
		// 點數
		Point int `json:"point,omitempty"`
		// 使用者電子郵件
		Email string `json:"email,omitempty"`
		// 角色ID
		RoleID string `json:"role_id,omitempty"`
		// 角色名稱
		RoleName string `json:"role_name,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"users"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`

	// 使用者名稱
	UserName string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name string `json:"name,omitempty"`
	// 使用者電話
	PhoneNumber string `json:"phone_number,omitempty"`
	// 是否啟用
	Active bool `json:"active,omitempty"`
	// 點數
	Point int `json:"point,omitempty"`
	// 使用者電子郵件
	Email string `json:"email,omitempty"`
	// 角色ID
	RoleID string `json:"role_id,omitempty"`
	// 角色名稱
	RoleName string `json:"role_name,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used to update achieves
type Update struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`

	// 使用者名稱
	UserName *string `json:"user_name,omitempty"`
	// 使用者中文名稱
	Name *string `json:"name,omitempty"`
	// 使用者密碼
	Password string `json:"password,omitempty"`
	// 使用者電話
	PhoneNumber *string `json:"phone_number,omitempty"`
	// 是否啟用
	Active *bool `json:"active,omitempty"`
	// 點數
	Point *int `json:"point,omitempty"`
	// 使用者電子郵件
	Email *string `json:"email,omitempty"`
	// 角色ID
	RoleID *string `json:"role_id,omitempty"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty"`
}

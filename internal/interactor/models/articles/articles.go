package articles

import (
	"ai_writer/internal/interactor/models/page"
	"ai_writer/internal/interactor/models/section"
)

// Create struct is used to create article
type Create struct {
	// 文章設定
	// 內容形式
	ContentType string `json:"content_type,omitempty" binding:"required" validate:"required"`
	// 論壇
	Forum string `json:"forum,omitempty" binding:"required" validate:"required"`
	// 版位
	Board string `json:"board,omitempty" binding:"required" validate:"required"`
	// 文章類型
	Type string `json:"type,omitempty" binding:"required" validate:"required"`
	// 文章風格
	Style string `json:"style,omitempty"`
	// 業配程度
	Sponsorship string `json:"sponsorship,omitempty"`
	// 文章標題
	Title string `json:"title,omitempty" binding:"required" validate:"required"`
	// 文章字數
	WordLimit int `json:"word_limit,omitempty" binding:"required" validate:"required"`
	// 關鍵訊息
	KeyMessage string `json:"key_message,omitempty"`
	// 口碑切角
	Story string `json:"story,omitempty"`

	// 人物設定
	// 人物性別
	Gender string `json:"gender,omitempty" binding:"required" validate:"required"`
	// 人物年齡
	Age []int `json:"age,omitempty" binding:"required" validate:"required"`
	// 人物年齡 (後端轉換用)
	Ages string `json:"ages,omitempty" swaggerignore:"true"`
	// 人物年齡(起)
	FromAge int `json:"from_age,omitempty" swaggerignore:"true"`
	// 人物年齡(起)
	ToAge int `json:"to_age,omitempty" swaggerignore:"true"`
	// 人物特質
	CharacterTrait string `json:"character_trait,omitempty"`
	// 人物細部描述
	CharacterRemarks string `json:"character_remarks,omitempty"`

	// 產品資訊
	// 產品名稱
	ProductName string `json:"product_name,omitempty" binding:"required" validate:"required"`
	// 產品特性
	ProductFeature string `json:"product_feature,omitempty" binding:"required" validate:"required"`
	// 產品亮點
	ProductHighlights string `json:"product_highlights,omitempty" binding:"required" validate:"required"`
	// 是否有競品
	HasCompetitive bool `json:"has_comparative"`
	// 生成文章
	AIArticle string `json:"ai_article,omitempty" swaggerignore:"true"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 表ID
	ID string `json:"id,omitempty" binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 內容形式
	ContentType *string `json:"content_type,omitempty" form:"content_type"`
	// 論壇
	Forum *string `json:"forum,omitempty" form:"forum"`
	// 版位
	Board *string `json:"board,omitempty" form:"board"`
	// 文章類型
	Type *string `json:"type,omitempty" form:"type"`
	// 文章風格
	Style *string `json:"style,omitempty" form:"style"`
	// 業配程度
	Sponsorship *string `json:"sponsorship,omitempty" form:"sponsorship"`
	// 文章標題
	Title *string `json:"title,omitempty" form:"title"`
	// 文章字數
	WordLimit *int `json:"word_limit,omitempty" form:"word_limit"`
	// 關鍵訊息
	KeyMessage *string `json:"key_message,omitempty" form:"key_message"`
	// 口碑切角
	Story *string `json:"story,omitempty" form:"story"`
	// 人物性別
	Gender *string `json:"gender,omitempty" form:"gender"`
	// 人物年齡
	Age *[]int `json:"age,omitempty" form:"age"`
	// 人物特質
	CharacterTrait *string `json:"character_trait,omitempty" form:"character_trait"`
	// 人物細部描述
	CharacterRemarks *string `json:"character_remarks,omitempty" form:"character_remarks"`
	// 產品名稱
	ProductName *string `json:"product_name,omitempty" form:"product_name"`
	// 產品特性
	ProductFeature *string `json:"product_feature,omitempty" form:"product_feature"`
	// 產品亮點
	ProductHighlights *string `json:"product_highlights,omitempty" form:"product_highlights"`
	// 是否有競品
	HasCompetitive *bool `json:"has_comparative,omitempty" form:"has_comparative"`
	// 生成文章
	AIArticle *string `json:"ai_article,omitempty" form:"ai_article"`
	// 修改後的文章
	ModifyArticle *string `json:"modify_article,omitempty" form:"modify_article"`
	// 文章分數
	Rating *int `json:"rating,omitempty" form:"rating"`
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
	Articles []*struct {
		// 表ID
		ID string `json:"id,omitempty"`
		// 文章標題
		Title string `json:"title,omitempty"`
		// 創建者
		CreatedBy string `json:"created_by,omitempty"`
		// 更新者
		UpdatedBy string `json:"updated_by,omitempty"`
		// 時間戳記
		section.TimeAt
	} `json:"articles"`
	// 分頁返回結構檔
	page.Total
}

// Single return structure file
type Single struct {
	// 表ID
	ID string `json:"id,omitempty"`
	// 文章標題
	Title string `json:"title,omitempty"`
	// 生成文章
	AIArticle string `json:"ai_article,omitempty"`
	// 創建者
	CreatedBy string `json:"created_by,omitempty"`
	// 更新者
	UpdatedBy string `json:"updated_by,omitempty"`
	// 時間戳記
	section.TimeAt
}

// Update struct is used end_date update achieves
type Update struct {
	// 表ID
	ID string `json:"id,omitempty"  binding:"omitempty,uuid4" validate:"omitempty,uuid4" swaggerignore:"true"`
	// 修改後的文章
	ModifyArticle *string `json:"modify_article,omitempty"`
	// 文章分數
	Rating *int `json:"rating,omitempty" binding:"omitempty,gte=0" validate:"omitempty,gte=0"`
	// 更新者
	UpdatedBy *string `json:"updated_by,omitempty" binding:"required,uuid4" validate:"required,uuid4" swaggerignore:"true"`
}

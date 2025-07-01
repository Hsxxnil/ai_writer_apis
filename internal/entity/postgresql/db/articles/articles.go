package articles

import (
	"ai_writer/internal/interactor/models/special"
)

// Table struct is article database table struct
type Table struct {
	// 內容形式
	ContentType string `gorm:"<-:create;column:content_type;type:text;not null;" json:"content_type"`
	// 論壇
	Forum string `gorm:"<-:create;column:forum;type:text;not null;" json:"forum"`
	// 版位
	Board string `gorm:"<-:create;column:board;type:text;not null;" json:"board"`
	// 文章類型
	Type string `gorm:"<-:create;column:type;type:text;not null;" json:"type"`
	// 文章風格
	Style string `gorm:"<-:create;column:style;type:text;not null;" json:"style"`
	// 業配程度
	Sponsorship string `gorm:"<-:create;column:sponsorship;type:text;not null;" json:"sponsorship"`
	// 文章標題
	Title string `gorm:"<-:create;column:title;type:text;not null;" json:"title"`
	// 文章字數
	WordLimit int `gorm:"<-:create;column:word_limit;type:int;not null;" json:"word_limit"`
	// 關鍵訊息
	KeyMessage string `gorm:"<-:create;column:key_message;type:text;not null;" json:"key_message"`
	// 口碑切角
	Story string `gorm:"<-:create;column:story;type:text;" json:"story"`
	// 人物性別
	Gender string `gorm:"<-:create;column:gender;type:text;not null;" json:"gender"`
	// 人物年齡(起)
	FromAge int `gorm:"<-:create;column:from_age;type:int;not null;" json:"from_age"`
	// 人物年齡(迄)
	ToAge int `gorm:"<-:create;column:to_age;type:int;not null;" json:"to_age"`
	// 人物特質
	CharacterTrait string `gorm:"<-:create;column:character_trait;type:text;not null;" json:"character_trait"`
	// 人物細部描述
	CharacterRemarks string `gorm:"<-:create;column:character_remarks;type:text;not null;" json:"character_remarks"`
	// 產品名稱
	ProductName string `gorm:"<-:create;column:product_name;type:text;not null;" json:"product_name"`
	// 產品特性
	ProductFeature string `gorm:"<-:create;column:product_feature;type:text;not null;" json:"product_feature"`
	// 產品亮點
	ProductHighlights string `gorm:"<-:create;column:product_highlights;type:text;not null;" json:"product_highlights"`
	// 是否有競品
	HasCompetitive *bool `gorm:"<-:create;column:has_comparative;type:boolean;default:false;not null;" json:"has_comparative"`
	// 生成文章
	AIArticle string `gorm:"<-:create;column:ai_article;type:text;" json:"ai_article"`
	// 修改後的文章
	ModifyArticle string `gorm:"<-:update;column:modify_article;type:text;" json:"modify_article"`
	// 文章分數
	Rating int `gorm:"<-:update;column:rating;type:int;" json:"rating"`
	// 引入後端專用
	special.Table
}

// Base struct is corresponding to article table structure file
type Base struct {
	// 內容形式
	ContentType *string `json:"content_type,omitempty"`
	// 論壇
	Forum *string `json:"forum,omitempty"`
	// 版位
	Board *string `json:"board,omitempty"`
	// 文章類型
	Type *string `json:"type,omitempty"`
	// 文章風格
	Style *string `json:"style,omitempty"`
	// 業配程度
	Sponsorship *string `json:"sponsorship,omitempty"`
	// 文章標題
	Title *string `json:"title,omitempty"`
	// 文章字數
	WordLimit *int `json:"word_limit,omitempty"`
	// 關鍵訊息
	KeyMessage *string `json:"key_message,omitempty"`
	// 口碑切角
	Story *string `json:"story,omitempty"`
	// 人物性別
	Gender *string `json:"gender,omitempty"`
	// 人物年齡(起)
	FromAge *int `json:"from_age,omitempty"`
	// 人物年齡(迄)
	ToAge *int `json:"to_age,omitempty"`
	// 人物特質
	CharacterTrait *string `json:"character_trait,omitempty"`
	// 人物細部描述
	CharacterRemarks *string `json:"character_remarks,omitempty"`
	// 產品名稱
	ProductName *string `json:"product_name,omitempty"`
	// 產品特性
	ProductFeature *string `json:"product_feature,omitempty"`
	// 產品亮點
	ProductHighlights *string `json:"product_highlights,omitempty"`
	// 是否有競品
	HasCompetitive *bool `json:"has_comparative,omitempty"`
	// 生成文章
	AIArticle *string `json:"ai_article,omitempty"`
	// 修改後的文章
	ModifyArticle *string `json:"modify_article,omitempty"`
	// 文章分數
	Rating *int `json:"rating,omitempty"`
	// 引入後端專用
	special.Base
}

func (t *Table) TableName() string {
	return "articles"
}

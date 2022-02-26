package models

// Category 类目表
type Category struct {
	Id uint64 `gorm:"primaryKey"`
	Name string `gorm:"name"`
	ParentId uint64 `gorm:"parent_id"`
	Level uint `gorm:"level"`
	Sort uint `gorm:"sort"`
	Created string `gorm:"created"`
	Updated string `gorm:"updated"`
}

// WebCategoryCreateParam 后台系统 类目创建参数模型
type WebCategoryCreateParam struct {
	Name string `json:"name" binding:"required"`
	ParentId uint64 `json:"parentId" binding:"required,gt=0"`
	Level uint `json:"level" binding:"required,oneof=1 2 3"`
	Sort uint `json:"sort" binding:"required,gt=0"`
}

// WebCategoryDeleteParam 后台系统 类目删除参数模型
type WebCategoryDeleteParam struct {
	Id uint64 `form:"id" binding:"required,gt=0"`
}

// WebCategoryUpdateParam 后台系统 类目更新参数模型
type WebCategoryUpdateParam struct {
	Id uint64 `json:"id" binding:"required,gt=0"`
	Name string `json:"name" binding:"required"`
	Sort uint `json:"sort" binding:"required,gt=0"`
}

// WebCategoryQueryParam 后台系统，类目查询参数模型
type WebCategoryQueryParam struct {
	Page Page
	Id uint64 `json:"id" binding:"omitempty,gt=0"`
	Name string `json:"name" binding:"omitempty,gt=0"`
	ParentId uint64 `json:"parentId" binding:"omitempty,gt=0"`
	Level uint `json:"level" binding:"omitempty,gt=0"`
}

// WebCategoryList 后台系统，类目列表传输模型
type WebCategoryList struct {
	Id uint64 `json:"id"`
	Name string `json:"name"`
	ParentId uint64 `josn:"parentId"`
	Level uint `json:"level"`
	Sort uint `json:"sort"`
}

// 后台系统， 类目选项传输模型
type WebCategoryOption struct {
	Value uint64 `json:"value"`
	Label string `json:"label"`
	Children []WebCategoryOption `json:"children"`
}

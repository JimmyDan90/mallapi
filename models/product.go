package models

// Product 商品数据库映射模型
type Product struct {
	Id uint64 `gorm:"primaryKey" json:"id"`
	CategoryId uint64 `gorm:"category_id" json:"categoryId"`
	Title string `gorm:"title" json:"title"`
	Description string `gorm:"description" json:"description"`
	Price float64 `gorm:"price" json:"price"`
	Amount uint `gorm:"amount" json:"amount"`
	Sales uint `gorm:"sales" json:"sales"`
	MainImage string `gorm:"main_image" json:"mainImage"`
	Delivery string `gorm:"delivery" json:"delivery"`
	Assurance string `gorm:"assurance" json:"assurance"`
	Name string `gorm:"name" json:"name"`
	Weight float32 `gorm:"weight" json:"weight"`
	Brand string `gorm:"brand" json:"brand"`
	Origin string `gorm:"origin" json:"origin"`
	ShelfLife uint `gorm:"shelf_life" json:"shelfLife"`
	NetWeight float32 `gorm:"net_weight" json:"netWeight"`
	UseWay string `gorm:"use_way" json:"useWay"`
	PackingWay string `gorm:"packing_way" json:"packingWay"`
	StorageConditions string `gorm:"storage_conditions" json:"storageConditions"`
	DetailImage string `gorm:"detail_image" json:"detailImage"`
	Status uint `gorm:"status" json:"status"`
	Created string `gorm:"created" json:"created"`
	Updated string `gorm:"updated" json:"updated"`
}

// WebProductCreateParam 后台系统 商品创建参数模型
type WebProductCreateParam struct {
	CategoryId uint64 `json:"categoryId" binding:"required,gt=0"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description"`
	Price float64 `json:"price" binding:"required,gt=0"`
	Amount uint `json:"amount" binding:"required,gt=0"`
	MainImage string `json:"mainImage" binding:"required"`
	Delivery string `json:"delivery" binding:"required"`
	Assurance string `json:"assurance" binding:"required"`
	Name string `json:"name" binding:"required"`
	Weight float32 `json:"weight" binding:"required"`
	Brand string `json:"brand"`
	Origin string `json:"origin"`
	ShelfLife uint `json:"shelfLife" binding:"required,gt=0"`
	NetWeight float32 `json:"netWeight" binding:"required,gt=0"`
	UseWay string `json:"useWay"`
	PackingWay string `json:"packingWay"`
	StorageConditions string `json:"storageConditions"`
	DetailImage string `json:"detailImage" binding:"required"`
	Status uint `json:"status" binding:"required,oneof=1 2"`
}

// WebProductIdParam 后台系统 ID参数
type WebProductIdParam struct {
	Id uint64 `json:"id" binding:"required,gt=0"`
}

// WebProductUpdateParam 后台系统商品更新参数模型
type WebProductUpdateParam struct {
	WebProductIdParam
	WebProductCreateParam
}

// WebProductStatusUpdateParam 后台系统商品状态更新模型
type WebProductStatusUpdateParam struct {
	WebProductIdParam
	Status uint `json:"status" binding:"required,oneof=1 2"`
}

// WebProductListParam 后台系统列表参数查询模型
type WebProductListParam struct {
	Page Page
	Id uint64 `json:"id" binding:"omitempty,gt=0"`
	CategoryId uint64 `json:"categoryId" binding:"omitempty,gt=0"`
	Title string `json:"title" binding:"omitempty"`
	Status uint `json:"status" binding:"omitempty,oneof=1 2"`
}

// WebProductInfo 后台系统商品信息传输模型
type WebProductInfo struct {
	Id uint64 `json:"id"`
	Category uint64 `json:"categoryId"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Amount uint `json:"amount"`
	MainImage string `json:"mainImage"`
	Delivery string `json:"delivery"`
	Assurance string `json:"assurance"`
	Name string `json:"name"`
	Weight float32 `json:"weight"`
	Brand string `json:"brand"`
	Origin string `json:"origin"`
	ShelfLife uint `json:"shelfLife"`
	NetWeight float32 `json:"netWeight"`
	UseWay string `json:"useWay"`
	PackingWay string `json:"packingWay"`
	StorageConditions string `json:"storageConditions"`
	DetailImage string `json:"detailImage"`
}

// WebProductItem 后台系统商品传输模型
type WebProductItem struct {
	Id uint64 `json:"id"`
	Title string `json:"title"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	MainImage string `json:"mainImage"`
}

// WebProductList 后台系统商品传输模型
type WebProductList struct {
	WebProductItem
	Sales int `json:"sales"`
	Status int `json:"status"`
	Created string `json:"created"`
}

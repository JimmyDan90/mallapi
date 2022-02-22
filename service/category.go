package service

import (
	"mallapi/common"
	"mallapi/global"
	"mallapi/models"
)

type WebCategoryService struct {
}

// Create 后台管理前端创建类目
func (c *WebCategoryService) Create(param models.WebCategoryCreateParam) uint64 {
	var category models.Category
	result := global.Db.Where("name = ?", param.Name).First(&category)
	if result.RowsAffected > 0 {
		return category.Id
	}
	category = models.Category{
		Name: param.Name,
		ParentId: param.ParentId,
		Level: param.Level,
		Sort: param.Sort,
		Created: common.NowTime(),
		Updated: common.NowTime(),
	}
	global.Db.Create(&category)
	return category.Id
}

// Delete 后台系统，删除类目
func (c *WebCategoryService) Delete(param models.WebCategoryDeleteParam) int64 {
	var pid2, pid3 models.Category
	global.Db.Where("parent_id = ?", param.Id).First(&pid2)
	global.Db.Where("parent_id = ?", pid2.Id).First(&pid3)
	return global.Db.Delete(&models.Category{}, []uint64{param.Id, pid2.Id, pid3.Id}).RowsAffected
}

// Update 后台系统，更新类目
func (c *WebCategoryService) Update(param models.WebCategoryUpdateParam) int64 {
	category := models.Category{
		Id: param.Id,
		Name: param.Name,
		Sort: param.Sort,
		Updated: common.NowTime(),
	}
	return global.Db.Model(&category).Updates(category).RowsAffected
}

// GetList 后台系统，获取类目列表
func (c *WebCategoryService) GetList(param models.WebCategoryQueryParam) ([]models.WebCategoryList, int64) {
	categoryList := make([]models.WebCategoryList, 0)
	query := &models.Category{
		Id: param.Id,
		Name: param.Name,
		Level: param.Level,
		ParentId: param.ParentId,
	}
	rows := common.RestPage(param.Page, "category", query, &categoryList, &[]models.Category{})
	return categoryList, rows
}

// GetOption 后台系统，获取类目选项
func (c *WebCategoryService) GetOption() (option []models.WebCategoryOption) {
	selectList := make([]models.WebCategoryList, 0)
	global.Db.Table("category").Find(&selectList)
	return getTreeOptions(1, selectList)
}

// 获取树形结构的选项
func getTreeOptions(id uint64, cateList []models.WebCategoryList) (option []models.WebCategoryOption) {
	optionList := make([]models.WebCategoryOption, 0)
	for _, opt := range cateList {
		if opt.ParentId == id && (opt.Level == 1 || opt.Level == 2) {
			option := models.WebCategoryOption{
				Value: opt.Id,
				Label: opt.Name,
				Children: getTreeOptions(opt.Id, cateList),
			}
			if opt.Level == 2 {option.Children = nil}
			optionList = append(optionList, option)
		}
	}
	return optionList
}

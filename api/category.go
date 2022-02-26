package api

import (
	"github.com/gin-gonic/gin"
	"mallapi/models"
	"mallapi/response"
	"mallapi/service"
)

var webCategory service.WebCategoryService

// WebCreateCategory 后台系统创建类目
// @Summary web后台系统创建类目接口
// @Description 需要传token
// @Accept application/json
// @Produce application/json
// @Param object body models.WebCategoryCreateParam true "类目创建参数"
// @Success 200 {object} string "{"message": "创建成功", count, c}"
// @Router /category/create [post]
func WebCreateCategory(c *gin.Context)  {
	var param models.WebCategoryCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Create(param); count > 0 {
		response.Success("创建成功", count ,c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteCategory 后台系统删除类目
func WebDeleteCategory(c *gin.Context)  {
	var param models.WebCategoryDeleteParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateCategory 后台系统更新类目
func WebUpdateCategory(c *gin.Context) {
	var param models.WebCategoryUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webCategory.Update(param); count > 0 {
		response.Success("更新成功", count ,c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetCategoryList 后台系统获取类目列表
func WebGetCategoryList(c *gin.Context) {
	var param models.WebCategoryQueryParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	cateList, rows := webCategory.GetList(param)
	response.SuccessPage("查询成功", cateList, rows, c)
}

// WebGetCategoryOption 后台系统，获取类目选项
func WebGetCategoryOption(c *gin.Context) {
	option := webCategory.GetOption()
	response.Success("查询成功", option, c)
}

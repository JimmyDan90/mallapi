package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mallapi/models"
	"mallapi/response"
	"mallapi/service"
)

var webProduct service.WebProductService

// WebCreateProduct 后台系统创建商品
func WebCreateProduct(c *gin.Context) {
	var param models.WebProductCreateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webProduct.Create(param); count > 0 {
		response.Success("创建成功", count ,c)
		return
	}
	response.Failed("创建失败", c)
}

// WebDeleteProduct 后台系统删除商品
func WebDeleteProduct(c *gin.Context) {
	var param models.WebProductIdParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webProduct.Delete(param); count > 0 {
		response.Success("删除成功", count, c)
		return
	}
	response.Failed("删除失败", c)
}

// WebUpdateProduct 后台系统，更新商品
func WebUpdateProduct(c *gin.Context) {
	var param models.WebProductUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		fmt.Println(err)
		response.Failed("请求参数无效", c)
		return
	}
	if count := webProduct.Update(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebUpdateProductStatus 后台系统 更新商品状态
func WebUpdateProductStatus(c *gin.Context) {
	var param models.WebProductStatusUpdateParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	if count := webProduct.UpdateStatus(param); count > 0 {
		response.Success("更新成功", count, c)
		return
	}
	response.Failed("更新失败", c)
}

// WebGetProductInfo 后台系统获取商品信息
func WebGetProductInfo(c *gin.Context) {
	var param models.WebProductIdParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productInfo := webProduct.GetInfo(param)
	response.Success("查询成功", productInfo, c)
}

// WebGetProductList 后台系统获取商品列表
func WebGetProductList(c *gin.Context) {
	var param models.WebProductListParam
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}
	productList, rows := webProduct.GetList(param)
	response.SuccessPage("查询成功", productList, rows, c)
}
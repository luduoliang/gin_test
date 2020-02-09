package api

import (
	"gin_test/models"
	"gin_test/common"
	"github.com/gin-gonic/gin"
)

//资讯分类列表，参数：
func CategoryList(c *gin.Context) {
	list := models.GetAllCategoryList()
	c.JSON(200, common.JsonData(true, list, "操作成功"))
}
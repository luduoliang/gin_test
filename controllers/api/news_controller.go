package api

import (
	"gin_test/models"
	"gin_test/common"
	"github.com/gin-gonic/gin"
	"strconv"
)

//资讯列表，参数：page,per_page,category_id(可传)
func NewsList(c *gin.Context) {
	page_temp := c.Request.FormValue("page")
	page, _ := strconv.Atoi(page_temp)
	per_page_temp := c.Request.FormValue("per_page")
	per_page, _ := strconv.Atoi(per_page_temp)
	category_id_temp := c.Request.FormValue("category_id")
	category_id, _ := strconv.Atoi(category_id_temp)

	list := models.GetNewsList(category_id, page, per_page)
	c.JSON(200, common.JsonData(true, list, "操作成功"))
}

//资讯详情，参数：id
func NewsDetail(c *gin.Context) {
	id_temp := c.Request.FormValue("id")
	id, _ := strconv.Atoi(id_temp)


	info := models.GetNewsPreload(id)
	c.JSON(200, common.JsonData(true, info, "操作成功"))
}
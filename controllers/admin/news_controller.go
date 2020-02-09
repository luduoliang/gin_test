package admin

import (
	"gin_test/models"
	"gin_test/library"
	"gin_test/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//列表
func NewsList(c *gin.Context){
	page_temp := c.Request.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetNewsListPaginator(page)
	paginator := library.NewPaginator(c.Request, models.PAGE_LIMIT, nums)

	c.HTML(http.StatusOK, "admin/news/list.tmpl", gin.H{
		"list": list,
		"Page": paginator,
	})
}

//添加
func NewsAdd(c *gin.Context) {
	if(c.Request.Method == "POST"){
		category_id_temp :=  c.Request.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title :=  c.Request.FormValue("title")
		description :=  c.Request.FormValue("description")
		pic :=  c.Request.FormValue("pic")
		content :=  c.Request.FormValue("content")
		weight_temp := c.Request.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info := models.News{}
		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.CreateNews(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()

		c.HTML(http.StatusOK, "admin/news/add.tmpl", gin.H{
			"category": category,
		})
	}

}

//编辑
func NewsEdit(c *gin.Context) {
	id_temp := c.Request.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetNews(id)
	if(c.Request.Method == "POST"){
		category_id_temp :=  c.Request.FormValue("category_id")
		category_id, _ := strconv.Atoi(category_id_temp)
		title :=  c.Request.FormValue("title")
		description :=  c.Request.FormValue("description")
		pic :=  c.Request.FormValue("pic")
		content :=  c.Request.FormValue("content")
		weight_temp := c.Request.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.CategoryId = category_id
		info.Title = title
		info.Description = description
		info.Pic = pic
		info.Content = content
		info.Weight = weight

		info = models.UpdateNews(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		category := models.GetAllCategoryList()

		c.HTML(http.StatusOK, "admin/news/edit.tmpl", gin.H{
			"category": category,
			"info": info,
		})
	}
}


//删除
func NewsDelete(c *gin.Context) {
	if(c.Request.Method == "POST"){
		id_temp := c.Request.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteNews(id)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
}

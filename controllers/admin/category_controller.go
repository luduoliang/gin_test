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
func CategoryList(c *gin.Context){
	page_temp := c.Request.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetCategoryListPaginator(page)
	paginator := library.NewPaginator(c.Request, models.PAGE_LIMIT, nums)

	c.HTML(http.StatusOK, "admin/category/list.tmpl", gin.H{
		"list": list,
		"Page": paginator,
	})
}

//添加
func CategoryAdd(c *gin.Context) {
	if(c.Request.Method == "POST"){
		name := c.Request.FormValue("name")
		weight_temp := c.Request.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)
		info := models.Category{}
		info.Name = name
		info.Weight = weight
		info = models.CreateCategory(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		c.HTML(http.StatusOK, "admin/category/add.tmpl", nil)
	}

}

//编辑
func CategoryEdit(c *gin.Context) {
	id_temp := c.Request.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetCategory(id)
	if(c.Request.Method == "POST"){
		name := c.Request.FormValue("name")
		weight_temp := c.Request.FormValue("weight")
		weight, _ := strconv.Atoi(weight_temp)

		info.Name = name
		info.Weight = weight
		info = models.UpdateCategory(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		c.HTML(http.StatusOK, "admin/category/edit.tmpl", gin.H{
			"info": info,
		})
	}
}


//删除
func CategoryDelete(c *gin.Context) {
	if(c.Request.Method == "POST"){
		id_temp := c.Request.FormValue("id")
		id, _ := strconv.Atoi(id_temp)
		info := models.DeleteCategory(id)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
}


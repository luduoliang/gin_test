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
func AdminList(c *gin.Context){
	page_temp := c.Request.FormValue("p")
	page, _ := strconv.Atoi(page_temp)
	list, nums := models.GetAdminListPaginator(page)
	paginator := library.NewPaginator(c.Request, models.PAGE_LIMIT, nums)

	c.HTML(http.StatusOK, "admin/admin/list.tmpl", gin.H{
		"list": list,
		"Page": paginator,
	})
}

//添加
func AdminAdd(c *gin.Context) {
	if(c.Request.Method == "POST"){
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		info := models.Admin{}
		info.Username = username
		info.Password = password
		info = models.CreateAdmin(info)

		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		c.HTML(http.StatusOK, "admin/admin/add.tmpl", nil)
	}

}

//编辑
func AdminEdit(c *gin.Context) {
	id_temp := c.Request.FormValue("id")
	id, _ := strconv.Atoi(id_temp)
	info := models.GetAdmin(id)
	if(c.Request.Method == "POST"){
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("password")

		info.Username = username
		info.Password = password
		info = models.UpdateAdmin(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}else{
		c.HTML(http.StatusOK, "admin/admin/edit.tmpl", gin.H{
			"info": info,
		})
	}
}


//删除
func AdminDelete(c *gin.Context) {
	if(c.Request.Method == "POST"){
		id_temp := c.Request.FormValue("id")
		id, _ := strconv.Atoi(id_temp)

		info := models.DeleteAdmin(id)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
}

//禁用/恢复
func AdminForbid(c *gin.Context) {
	if(c.Request.Method == "POST"){
		id_temp := c.Request.FormValue("id")
		id, _ := strconv.Atoi(id_temp)
		info := models.GetAdmin(id)

		if(info.Status == 1){
			info.Status = 2
		}else if(info.Status == 2){
			info.Status = 1
		}

		info = models.UpdateAdminInfo(info)
		c.JSON(200, common.JsonData(true, info, "操作成功"))
	}
}


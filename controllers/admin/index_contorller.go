package admin

import (
	"strconv"
	"time"
	"os"
	"gin_test/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_test/models"
)

//首页
func Index(c *gin.Context){

	admin_info_temp := common.GetSession(c, "admin_info")
	var admin_info models.Admin
	if(admin_info_temp != nil){
		admin_info = admin_info_temp.(models.Admin)
	}

	c.HTML(http.StatusOK, "admin/index/index.tmpl", gin.H{
		"admin_info": admin_info,
	})
}

//首页信息
func Info(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/info.tmpl", nil)
}

//上传文件
func UploadFile(c *gin.Context) {
	file, form_err := c.FormFile("pic")
	if form_err != nil {
		c.JSON(200, common.JsonData(false, "", form_err.Error()))
		return
	}

	fname := file.Filename

	year := strconv.Itoa(time.Now().Year())
	month := time.Now().Month().String()
	day := strconv.Itoa(time.Now().Day())
	//创建目录
	dirpath := "static/upload/"+year+"/"+month+"/"+day
	dir_err := os.MkdirAll(dirpath, os.ModePerm)
	if dir_err != nil {
		c.JSON(200, common.JsonData(false, "", dir_err.Error()))
		return
	}

	//上传文件相对路劲
	file_path := dirpath + "/" + fname

	out, open_err := os.OpenFile(file_path, os.O_WRONLY|os.O_CREATE, 0666)
	defer out.Close()

	if open_err != nil {
		c.JSON(200, common.JsonData(false, "", open_err.Error()))
		return
	}

	copy_err := c.SaveUploadedFile(file, file_path)
	if(copy_err != nil){
		c.JSON(200, common.JsonData(false, "", copy_err.Error()))
		return
	}

	scheme := "http://"
	if c.Request.TLS != nil {
		scheme = "https://"
	}

	all_file_path := scheme + c.Request.Host + "/" + file_path
	upload_info := make(map[string]string)
	upload_info["pic_url"] = all_file_path

	c.JSON(200, common.JsonData(true, upload_info, "上传成功"))
}



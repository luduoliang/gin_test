package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gin_test/common"
)


//验证用户是否登录
func CheckAdminLogin() gin.HandlerFunc {
	return func (c *gin.Context) {
		admin_info := common.GetSession(c, "admin_info")
		if admin_info == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}

		c.Next()
	}
}

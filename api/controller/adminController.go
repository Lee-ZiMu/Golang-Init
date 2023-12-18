/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/7 17:28
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"hy_heymate/api/service"
	token "hy_heymate/common/jwt"
	"hy_heymate/common/result"
)

func AdminLogin(c *gin.Context) {
	AppleID := c.Query("AppleID")
	h := service.NewLoginHandler(c.Request.Context())
	token, err := h.Login(AppleID)
	result.GinResult(c, token, err)
}

func AdminLogout(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	token.CheckToken(c)
	h := service.NewLoginHandler(c.Request.Context())
	str, err := h.Logout(authorization)
	result.GinResult(c, str, err)
}

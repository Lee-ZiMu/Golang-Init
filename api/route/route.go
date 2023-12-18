/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/7 17:01
 **/

package route

import (
	"github.com/gin-gonic/gin"
	"hy_heymate/api/controller"
	token "hy_heymate/common/jwt"
	"hy_heymate/common/ws"
)

func LoadRouters(r *gin.Engine) {
	r.Use(Cors())

	r.GET("/ws", ws.HandleWebSocket)

	// 登录认证模块
	admin := r.Group("/client/admin")
	admin.POST("/login", controller.AdminLogin)   // 登录
	admin.POST("/logout", controller.AdminLogout) // 登出

	// 个人中心
	personal := r.Group("/client/personalInformation")
	personal.Use(token.CheckToken)
	personal.GET("/get", controller.GetPersonalInformation)        // 获取个人资料
	personal.POST("/update", controller.UpdatePersonalInformation) // 修改个人资料

}

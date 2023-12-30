/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/6 16:38
 **/

package main

import (
	"github.com/gin-gonic/gin"
	viperInit "github.com/spf13/viper"
	"hy_heymate/api/route"
	"hy_heymate/common/cron"
	"hy_heymate/common/logger"
	"hy_heymate/config/viper"
	"hy_heymate/database"
)

func main() {
	r := gin.Default()

	// 读取配置文件
	viper.Init("../config/config.yaml")

	// 初始化日志
	logger.Init()

	// 连接mysql数据库
	database.ConnectToMySQL()

	// 连接redis
	database.ConnectToRedis()

	route.LoadRouters(r)

	// 定时
	cron.Cron()

	r.Run(":" + viperInit.GetString("server.port"))

}

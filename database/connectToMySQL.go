/**
 * @Description: 连接数据库
 * @Author Lee
 * @Date 2023/12/7 9:08
 **/

package database

import (
	"fmt"
	viperInit "github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"hy_heymate/common/file"
	"hy_heymate/gen/query"
	"log"
	"os"
	"time"
)

var db = new(gorm.DB)

func ConnectToMySQL() {
	// 创建日志文件夹
	if err := file.MkdirAll(viperInit.GetString("mysql.logpath")); err != nil {
		panic(err)
	}

	logFile, err := os.Create(viperInit.GetString("mysql.logpath") + "db.log")
	if err != nil {
		panic(err)
	}

	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		viperInit.GetString("mysql.username"),
		viperInit.GetString("mysql.password"),
		viperInit.GetString("mysql.host"),
		viperInit.GetInt("mysql.port"),
		viperInit.GetString("mysql.dbname"),
		viperInit.GetString("mysql.conf"),
	)
	db, err = gorm.Open(mysql.Open(sqlStr), &gorm.Config{
		Logger: logger.New(
			log.New(logFile, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			}),
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		panic(fmt.Errorf("mysql连接失败，错误信息：%s", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viperInit.GetInt("mysql.maxIdleConns"))                                      // 最大连接数
	sqlDB.SetMaxOpenConns(viperInit.GetInt("mysql.maxOpenConns"))                                      // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Duration(viperInit.GetInt64("mysql.vonnMaxLifetime")) * time.Second) // 设置可以重用连接的最大时间量

	query.SetDefault(db)

}

func Get() *gorm.DB {
	return db
}

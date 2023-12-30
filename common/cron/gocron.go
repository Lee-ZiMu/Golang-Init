/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/30 10:50
 **/

package cron

import (
	"github.com/robfig/cron/v3"
	"hy_heymate/common/logger"
)

var gocron *cron.Cron

func Cron() {
	gocron = cron.New(cron.WithSeconds())

	// 每日0时重置任务
	_, err := gocron.AddFunc("0 0 0 * * *", func() {
		// 业务代码
		// ................
		// ................
	})

	if err != nil {
		logger.Errorf("Timed Reset Task Error: %v", err)
		return
	}

	// 启动定时任务
	gocron.Start()

}

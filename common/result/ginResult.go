/**
 * @Description: 返回数据
 * @Author Lee
 * @Date 2023/12/11 13:50
 **/

package result

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"hy_heymate/common/errType"
	"hy_heymate/common/logger"
	"net/http"
)

func GinResult(c *gin.Context, data interface{}, err error) {
	if err == nil {
		// 成功返回
		c.JSON(http.StatusOK, Success(data))
		return
	}

	// 打印错误日志到控制台
	log, _ := zap.NewDevelopment()
	log.Error(fmt.Sprintf("操作失败，错误信息：%s", err))

	// 错误返回
	errcode := errType.ServeError
	errmsg := errType.GetErrorMsg(errType.ServeError)

	causeErr := errors.Cause(err) // err类型
	e, ok := causeErr.(*errType.CodeError)
	if ok { // 自定义错误类型
		// 自定义CodeError
		errcode = e.GetErrCode()
		errmsg = e.GetErrMsg()
	}
	logger.Errorf("http-error: %v", err)
	c.JSON(http.StatusOK, Error(errcode, errmsg))

}

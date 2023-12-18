/**
 * @Description: 响应模板
 * @Author Lee
 * @Date 2023/12/11 13:31
 **/

package result

import (
	"hy_heymate/common/errType"
)

type ResponseTemplate struct {
	Code    string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *ResponseTemplate {
	return &ResponseTemplate{errType.Success, errType.GetErrorMsg(errType.Success), data}
}

func Error(errCode string, errMsg string) *ResponseTemplate {
	return &ResponseTemplate{errCode, errMsg, nil}
}

func ErrorWithData(errCode, errMsg string, data interface{}) *ResponseTemplate {
	return &ResponseTemplate{errCode, errMsg, data}
}

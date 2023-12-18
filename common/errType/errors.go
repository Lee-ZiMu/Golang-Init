/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/11 15:06
 **/

package errType

import "fmt"

type CodeError struct {
	errCode, errMsg string
}

func (e *CodeError) GetErrCode() string {
	return e.errCode
}

func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%s,ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode, errMsg string) *CodeError {
	e := &CodeError{errCode: errCode, errMsg: errMsg}
	return e
}

func NewErrCode(errCode string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: GetErrorMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: ServeError, errMsg: errMsg}
}

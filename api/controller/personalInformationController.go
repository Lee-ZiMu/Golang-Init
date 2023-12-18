/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/7 17:36
 **/

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"hy_heymate/api/service"
	"hy_heymate/common/errType"
	"hy_heymate/common/result"
	"hy_heymate/common/vo"
)

func GetPersonalInformation(c *gin.Context) {
	data, err := service.GetPersonalInformation(c, cast.ToInt64(c.GetString("userId")))
	result.GinResult(c, data, err)
}

func UpdatePersonalInformation(c *gin.Context) {
	var req vo.PersonalInformation
	err := c.ShouldBindJSON(&req)
	if err != nil {
		result.GinResult(c, nil, errors.Wrapf(errType.NewErrCode(errType.TypeConversion), "Conversion struct failed,error:%v", err))
		return
	}
	err = service.UpdatePersonalInformation(c, c.GetInt64("userId"), &req)
	result.GinResult(c, nil, err)
}

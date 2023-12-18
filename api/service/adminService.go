/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/7 17:47
 **/

package service

import (
	"context"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	tp "hy_heymate/common/const"
	"hy_heymate/common/errType"
	"hy_heymate/common/redisUtil"
	"hy_heymate/common/vo"
	"hy_heymate/gen/query"
	"strings"
)

type LoginHandler struct {
	ctx context.Context
}

func NewLoginHandler(ctx context.Context) *LoginHandler {
	return &LoginHandler{
		ctx: ctx,
	}
}

func (h *LoginHandler) Login(AppleID string) (*vo.ResSysUser, error) {
	// 查询user_id
	pi, pic := query.SysUser, query.SysUser.WithContext(h.ctx)
	first, err := pic.Where(pi.AppleID.Eq(AppleID)).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(errType.NewErrCode(errType.DatabaseSelectFailed), "Failed to select data,errormsg:%v", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 初始化

		// 添加至redis
		return nil, nil
	} else {
		s := uuid.NewV4().String()
		sysUser := &vo.ResSysUser{
			Token:          s,
			Isusefree:      first.Isusefree,
			Type:           first.Type,
			ExpirationTime: first.ExpirationTime,
		}
		// UserID添加至redis
		_, err := redisUtil.SetString(tp.TokenPrefix+s, cast.ToString(first.UserID))
		if err != nil {
			return nil, err
		}
		return sysUser, nil
	}
}

func (h *LoginHandler) Logout(tokenString string) (bool, error) {
	split := strings.Split(tokenString, "_")
	delString, err := redisUtil.DelString(split[1])
	return delString, err
}

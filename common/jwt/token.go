/**
 * @Description: token验证
 * @Author Lee
 * @Date 2023/12/15 10:35
 **/

package token

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	tp "hy_heymate/common/const"
	"hy_heymate/common/errType"
	"hy_heymate/common/logger"
	"hy_heymate/common/redisUtil"
	"hy_heymate/common/result"
)

func CheckToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		logger.Errorf("Token is empty")
		result.GinResult(c, nil, errors.Wrap(errType.NewErrCode(errType.TokenIsNil), "Token cannot be empty"))
		c.Abort()
		return
	}
	userId, err := GetUserIdByToken(token)
	if err != nil {
		result.GinResult(c, nil, err)
		c.Abort()
		return
	} else {
		c.Set("userId", userId)
		c.Next()
		return
	}
}

func GetUserIdByToken(tokenString string) (string, error) {
	exists, err := redisUtil.Exists(tp.TokenPrefix + tokenString)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", errors.Wrapf(errType.NewErrCode(errType.TokenExpires), "Token Expires,errormsg:%v", err)
	}
	userId, err := redisUtil.GetString(tp.TokenPrefix + tokenString)
	return userId, err
}

// jwt
/*
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"hy_heymate/common/errType"
	"hy_heymate/common/logger"
	"hy_heymate/common/result"
	"time"
)

type MyCustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func GetToken(userId int64) (string, error) {
	mySigningKey := []byte("AllYourBase")
	// Create the Claims
	m, _ := time.ParseDuration("-1m")
	d, _ := time.ParseDuration("24h")
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			NotBefore: time.Now().Add(m).Unix(), // 生效时间
			ExpiresAt: time.Now().Add(d).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		logger.ErrorE("Get Token Error", err)
		return "", errors.Wrap(errType.NewErrCode(errType.GetTokenError), "Get Token Error")
	}
	return signedString, err
}

func CheckToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		logger.Errorf("Token is empty")
		result.GinResult(c, nil, errors.Wrap(errType.NewErrCode(errType.TokenIsNil), "Token cannot be empty"))
		c.Abort()
		return
	}
	userId, err := ParseToken(token)
	if err != nil {
		result.GinResult(c, nil, err)
		c.Abort()
		return
	} else {
		c.Set("userId", userId)
		c.Next()
		return
	}
}

func ParseToken(tokenString string) (int64, error) {
	mySigningKey := []byte("AllYourBase")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		logger.ErrorE("Token illegal", err)
		return 0, errors.Wrap(errType.NewErrCode(errType.TokenIllegal), "Token illegal")
	}
	id := token.Claims.(*MyCustomClaims).UserId
	return id, nil
}

*/

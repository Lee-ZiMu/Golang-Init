/**
 * @Description: redis
 * @Author Lee
 * @Date 2023/12/8 11:23
 **/

package redisUtil

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"hy_heymate/common/errType"
	"hy_heymate/database"
)

func SetString(key string, value string) (string, error) {
	r := database.RedisClient.Get()
	res, err := redis.String(r.Do("SET", key, value))
	if err != nil {
		return res, errors.Wrapf(errType.NewErrCode(errType.RedisSaveFailed), "Redis Save failed,errormsg:%v", err)
	}
	return res, err
}

func SetStringWithTime(key string, value string, outTime int) (string, error) {
	r := database.RedisClient.Get()
	var res string
	var err error
	if outTime > 0 {
		res, err = redis.String(r.Do("SET", key, value, "EX", outTime))
	} else {
		res, err = redis.String(r.Do("SET", key, value))
	}
	if err != nil {
		return res, errors.Wrapf(errType.NewErrCode(errType.RedisSaveFailed), "Redis Save failed,errormsg:%v", err)
	}
	return res, err
}

func GetString(key string) (string, error) {
	r := database.RedisClient.Get()
	res, err := redis.String(r.Do("GET", key))
	if err != nil {
		return res, errors.Wrapf(errType.NewErrCode(errType.RedisSelectFailed), "Redis Select Failed,errormsg:%v", err)
	}
	return res, err
}

func DelString(key string) (bool, error) {
	r := database.RedisClient.Get()
	res, err := redis.Bool(r.Do("DEL", key))
	if err != nil {
		return res, errors.Wrapf(errType.NewErrCode(errType.RedisDeleteFailed), "Redis Delete Failed,errormsg:%v", err)
	}
	return res, err
}

func Exists(key string) (bool, error) {
	r := database.RedisClient.Get()
	res, err := redis.Bool(r.Do("EXISTS", key))
	if err != nil {
		return res, errors.Wrapf(errType.NewErrCode(errType.RedisDoesItExist), "Redis Does It Exist,errormsg:%v", err)
	}
	return res, err
}

/**
 * @Description: 错误信息
 * @Author Lee
 * @Date 2023/12/7 16:13
 **/

package errType

var (
	message map[string]string
)

func GetErrorMsg(code string) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(code string) bool {
	_, ok := message[code]
	return ok
}

func init() {
	message = make(map[string]string)
	message[Success] = "请求成功"
	message[ServeError] = "服务器异常"

	message[ParameterError] = "请求参数错误"
	message[TypeConversion] = "类型转换错误"

	message[TokenExpires] = "Token过期"
	message[NoPermission] = "没有权限"
	message[GetTokenError] = "生成Token失败"
	message[TokenIllegal] = "非法Token"
	message[TokenIsNil] = "Token不能为空"

	message[DBError] = "数据库错误"
	message[RecordNotFound] = "查询数据为空"
	message[DatabaseSaveFailed] = "数据库添加数据失败"
	message[DatabaseDeleteFailed] = "数据库删除数据失败"
	message[DatabaseUpdateFailed] = "数据库更改数据失败"
	message[DatabaseSelectFailed] = "数据库查询数据失败"

	message[RedisError] = "Redis错误"
	message[RedisRecordNotFound] = "Redis查询数据为空"
	message[RedisSaveFailed] = "Redis添加数据失败"
	message[RedisDeleteFailed] = "Redis删除数据失败"
	message[RedisSelectFailed] = "Redis查询数据失败"
	message[RedisDoesItExist] = "Redis是否存在数据失败"

}

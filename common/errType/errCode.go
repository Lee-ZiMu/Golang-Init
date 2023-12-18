/**
 * @Description: 错误类型
 * @Author Lee
 * @Date 2023/12/7 16:05
 **/

package errType

// 通用错误码
const (
	Success    = "000000" // 操作成功
	ServeError = "400000" // 服务端操作失败
)

// 请求参数处理
const (
	ParameterError = "100001" // 请求参数错误
	TypeConversion = "100002" // 类型转换错误
)

// 授权错误码
const (
	TokenExpires  = "200001" // token过期
	NoPermission  = "200002" // 没有权限
	GetTokenError = "200003" // 生成token失败
	TokenIllegal  = "200004" // 非法token
	TokenIsNil    = "200005" // token不能为空
)

// 数据库错误
const (
	DBError              = "300000" // 数据库错误
	RecordNotFound       = "300001" // 查询数据为空
	DatabaseSaveFailed   = "300002" // 数据库添加数据失败
	DatabaseDeleteFailed = "300003" // 数据库删除数据失败
	DatabaseUpdateFailed = "300004" // 数据库更改数据失败
	DatabaseSelectFailed = "300005" // 数据库查询数据失败

	RedisError          = "300010" // Redis错误
	RedisRecordNotFound = "300011" // Redis查询数据为空
	RedisSaveFailed     = "300012" // Redis添加数据失败
	RedisDeleteFailed   = "300013" // Redis删除数据失败
	RedisSelectFailed   = "300014" // Redis查询数据失败
	RedisDoesItExist    = "300015" // Redis是否存在数据失败
)

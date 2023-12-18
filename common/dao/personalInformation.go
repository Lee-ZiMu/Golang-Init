/**
 * @Description: 用户与数据库交互结构体
 * @Author Lee
 * @Date 2023/12/11 16:46
 **/

package dao

type PersonalInformation struct {
	UserId     int64       `json:"user_id"`
	FileId     int         `json:"file_id"`
	Nickname   string      `json:"nickname"`
	Sex        interface{} `json:"sex"`
	ReignTitle interface{} `json:"reign_title"`
	Address    interface{} `json:"address"`
	Work       interface{} `json:"work"`
	Maxim      interface{} `json:"maxim"`
}

/**
 * @Description: 个人信息
 * @Author Lee
 * @Date 2023/12/18 9:31
 **/

package vo

type PersonalInformation struct {
	UID        int64  `gorm:"column:uid;type:int(11);primaryKey;comment:uid" json:"uid"`                  // uid
	Path       string `gorm:"column:file_id;type:text;not null;comment:头像路径" json:"path"`                 // 头像路径
	Nickname   string `gorm:"column:nickname;type:varchar(32);not null;comment:昵称" json:"nickname"`       // 昵称
	Sex        int64  `gorm:"column:sex;type:int(1);not null;comment:性别（1男，2女）" json:"sex"`               // 性别（1男，2女）
	ReignTitle string `gorm:"column:reign_title;type:varchar(10);not null;comment:年号" json:"reign_title"` // 年号
	Address    string `gorm:"column:address;type:varchar(32);not null;comment:所在地" json:"address"`        // 所在地
	Work       string `gorm:"column:work;type:varchar(32);not null;comment:职业" json:"work"`               // 职业
	Maxim      string `gorm:"column:maxim;type:varchar(64);not null;comment:箴言" json:"maxim"`             // 箴言
}

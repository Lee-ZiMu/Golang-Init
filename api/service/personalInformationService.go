/**
 * @Description:
 * @Author Lee
 * @Date 2023/12/7 17:47
 **/

package service

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gen"
	"gorm.io/gorm"
	"hy_heymate/common/errType"
	"hy_heymate/common/logger"
	"hy_heymate/common/vo"
	"hy_heymate/gen/model"
	"hy_heymate/gen/query"
)

func GetPersonalInformation(c *gin.Context, userIdInt64 int64) (*vo.PersonalInformation, error) {
	pi, pic := query.PersonalInformation, query.PersonalInformation.WithContext(c.Request.Context())
	first, err := pic.Where(pi.UserID.Eq(userIdInt64)).First()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.Wrapf(errType.NewErrCode(errType.RecordNotFound), "The query data is empty,errormsg:%v", err)
	}
	if err != nil {
		return nil, errors.Wrapf(errType.NewErrCode(errType.DatabaseSelectFailed), "Failed to select data,error:%v", err)
	}
	req := &vo.PersonalInformation{
		UID: first.UID,
		//Path:       first.FileID,
		Nickname:   first.Nickname,
		Sex:        first.Sex,
		ReignTitle: first.ReignTitle,
		Address:    first.Address,
		Work:       first.Work,
		Maxim:      first.Maxim,
	}
	return req, nil

	/*
		// 预编译sql
		var results []vo.PersonalInformation
		database.Get().Table("personal_information pi").
			Select("pi.uid as UID, sf.path as Path, pi.nickname as Nickname, pi.sex as Sex, pi.reign_title as ReignTitle, pi.address as Address, pi.work as Work, pi.maxim as Maxim").
			Joins("left join sys_file sf on sf.file_id = pi.file_id where pi.user_id = ?", cast.ToString(userIdInt64)).
			Scan(&results)
		return &results[0], nil
	*/

	/*
		// 存在sql注入
		sql := "select " +
			"pi.uid as UID, " +
			"sf.path as Path, " +
			"pi.nickname as Nickname, " +
			"pi.sex as Sex, " +
			"pi.reign_title as ReignTitle, " +
			"pi.address as Address, " +
			"pi.work as Work, " +
			"pi.maxim as Maxim " +
			"from personal_information pi " +
			"left join sys_file sf on sf.file_id = pi.file_id " +
			"where pi.user_id = '" + cast.ToString(userIdInt64) + "'"
		var results []vo.PersonalInformation
		database.Get().Raw(sql).Scan(&results)
		return &results[0], nil

	*/

}

func UpdatePersonalInformation(c *gin.Context, userID int64, req *vo.PersonalInformation) error {
	// 修改数据
	pi, pic := query.PersonalInformation, query.PersonalInformation.WithContext(c.Request.Context())
	_, err := pic.Where(pi.UserID.Eq(userID)).Updates(req)
	if err != nil {
		return errors.Wrapf(errType.NewErrCode(errType.DatabaseUpdateFailed), "Failed to update data,error:%v", err)
	}
	return nil
}

func AddPersonalInformation(c *gin.Context, req *model.PersonalInformation) (*model.PersonalInformation, error) {
	// 插入数据库
	_, pic := query.PersonalInformation, query.PersonalInformation.WithContext(c.Request.Context())
	err := pic.Save(req)
	if err != nil {
		logger.ErrorE("Failed to save data", err)
		return req, errors.Wrap(errType.NewErrCode(errType.DatabaseSaveFailed), "Failed to save data")
	}
	return req, nil
}

func DeletePersonalInformation(c *gin.Context, userIdInt64 int64) (gen.ResultInfo, error) {
	pi, pic := query.PersonalInformation, query.PersonalInformation.WithContext(c.Request.Context())
	info, err := pic.Where(pi.UserID.Eq(userIdInt64)).Delete()
	if err != nil {
		logger.ErrorE("Failed to delete data", err)
		return info, errors.Wrap(errType.NewErrCode(errType.DatabaseDeleteFailed), "Failed to delete data")
	}
	return info, nil
}

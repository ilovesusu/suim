package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"gorm.io/gorm"
	"time"
)

// FriendRepo .
type FriendRepo struct {
	data *Data
	log  *log.Helper
}

// CreateFriend 添加好友
func (fr *FriendRepo) CreateFriend(_ context.Context, req *biz.CreateFriendReq) error {
	if err := applyFriend(fr.data.db, req.UserId, req.FriendId, req.VerifyMessage); err != nil {
		fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v申请用户id%v成为好友失败", req.UserId, req.FriendId), err)
		return err
	}
	return nil
}

// ListUserFriend 列表
func (fr *FriendRepo) ListUserFriend(_ context.Context, req *biz.ListFriendReq) ([]biz.ListUserFriendRsp, error) {
	friends := make([]biz.UserFriend, 0)
	if err := fr.data.db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_status = ? and delete_time is null", req.UserId, req.FriendStatus).
		Find(&friends).Error; err != nil {
		fr.log.Log(log.LevelError, "查询好友关系失败", err)
		return nil, v1.ErrorErrorFail("查询失败")
	}
	data := make([]biz.ListUserFriendRsp, 0)
	for _, v := range friends {
		user := biz.ListUserFriendRsp{}
		if err := fr.data.db.Model(&biz.UserInfo{}).Select("nickname", "sex", "avatar_url", "personal_sign").Where("id = ?", v.FriendId).
			First(&user).Error; err != nil {
			fr.log.Log(log.LevelError, fmt.Sprintf("查找id为%v的用户失败", v.FriendId), err)
			continue
		}
		data = append(data, biz.ListUserFriendRsp{
			Id:           v.FriendId,
			Nickname:     user.Nickname,
			Sex:          user.Sex,
			AvatarUrl:    user.AvatarUrl,
			PersonalSign: user.PersonalSign,
			FriendRemark: v.FriendRemark,
		})
	}
	return data, nil
}

// UpdateFriendRemark 修改好友备注
func (fr *FriendRepo) UpdateFriendRemark(_ context.Context, req *biz.UpdateFriendRemarkReq) error {
	status, err := searchFriendStatus(fr.data.db, req.UserId, req.FriendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return v1.ErrorErrorFail("好友已注销帐号")
		} else {
			fr.log.Log(log.LevelError, fmt.Sprintf("修改用户id为%v好友id为%v的备注为%v失败", req.UserId, req.FriendId, req.FriendRemark), err)
			return v1.ErrorErrorFail("修改好友备注失败")
		}
	}
	//已拒绝的好友和已删除的好友不能修改备注
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_REFUSE) || *status == int32(friend.FriendStatus_FRIEND_STATUS_DELETE) {
		return v1.ErrorErrorFail("该用户不是您的好友")
	}
	if err := fr.data.db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", req.UserId, req.FriendId).
		Updates(map[string]interface{}{"friend_remark": req.FriendRemark}).Error; err != nil {
		fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v修改好友id%v的备注失败", req.UserId, req.FriendId), err)
		return v1.ErrorErrorFail("修改失败")
	}
	return nil
}

// UpdateFriendStatus 修改好友状态
func (fr *FriendRepo) UpdateFriendStatus(_ context.Context, req *biz.UpdateFriendStatusReq) error {
	//通过好友请求
	if req.FriendStatus == int32(friend.FriendStatus_FRIEND_STATUS_VERIFYING) {
		if err := verifyFriend(fr.data.db, req.UserId, req.FriendId); err != nil {
			fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v通过用户id%v好友请求失败", req.UserId, req.FriendId), err)
			return err
		} else {
			return nil
		}
	}
	//拒绝好友请求
	if req.FriendStatus == int32(friend.FriendStatus_FRIEND_STATUS_REFUSE) {
		if err := refuseFriend(fr.data.db, req.UserId, req.FriendId); err != nil {
			fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v拒绝用户id%v申请好友失败", req.UserId, req.FriendId), err)
			return err
		} else {
			return nil
		}
	}
	//好友拉入黑名单
	if req.FriendStatus == int32(friend.FriendStatus_FRIEND_STATUS_BLACK) {
		if err := blackFriend(fr.data.db, req.UserId, req.FriendId); err != nil {
			fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v将用户id%v加入黑名单失败", req.UserId, req.FriendId), err)
			return err
		} else {
			return nil
		}
	}
	//删除好友
	if req.FriendStatus == int32(friend.FriendStatus_FRIEND_STATUS_DELETE) {
		if err := deleteFriend(fr.data.db, req.UserId, req.FriendId); err != nil {
			fr.log.Log(log.LevelError, fmt.Sprintf("用户id%v删除用户id%v好友关系失败", req.UserId, req.FriendId), err)
			return err
		} else {
			return nil
		}
	}
	return v1.ErrorErrorFail("无效操作")
}

//申请好友
func applyFriend(db *gorm.DB, userId, friendId int64, message string) error {
	db = db.Begin()
	var err error
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	status, err := searchFriendStatus(db, userId, friendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			var userFriend biz.UserFriend
			userFriend.CreateTime = time.Now()
			userFriend.UserId = userId
			userFriend.FriendId = friendId
			userFriend.FriendStatus = int32(friend.FriendStatus_FRIEND_STATUS_APPLYING)
			userFriend.ApplyMessage = message
			if err := db.Model(&biz.UserFriend{}).Create(&userFriend).Error; err != nil {
				return v1.ErrorErrorFail("操作失败")
			}
			userFriend.UserId = friendId
			userFriend.FriendId = userId
			userFriend.FriendStatus = int32(friend.FriendStatus_FRIEND_STATUS_VERIFYING)
			if err := db.Model(&biz.UserFriend{}).Create(&userFriend).Error; err != nil {
				return v1.ErrorErrorFail("操作失败")
			}
		} else {
			return v1.ErrorErrorFail("操作失败")
		}
	}
	//已拒绝的好友和已删除的好友重新申请
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_REFUSE) || *status == int32(friend.FriendStatus_FRIEND_STATUS_DELETE) {
		if err := db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", userId, friendId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_APPLYING, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
		if err := db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", friendId, userId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_VERIFYING, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
	} else {
		return v1.ErrorErrorFail("已是好友关系,无需重复添加")
	}
	return nil
}

//验证好友
func verifyFriend(db *gorm.DB, userId, friendId int64) error {
	var err error
	db = db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	status, err := searchFriendStatus(db, userId, friendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return v1.ErrorErrorFail("无效操作")
		} else {
			return v1.ErrorErrorFail("操作失败")
		}
	}
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_VERIFYING) {
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", userId, friendId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_PASS, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", friendId, userId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_PASS, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
	} else {
		return v1.ErrorErrorFail("无效操作")
	}
	return nil
}

//拒绝好友
func refuseFriend(db *gorm.DB, userId, friendId int64) error {
	var err error
	db = db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	status, err := searchFriendStatus(db, userId, friendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return v1.ErrorErrorFail("无效操作")
		} else {
			return v1.ErrorErrorFail("操作失败")
		}
	}
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_VERIFYING) {
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", userId, friendId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_REFUSE, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", friendId, userId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_REFUSE, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
	} else {
		return v1.ErrorErrorFail("无效操作")
	}
	return nil
}

//拉黑好友
func blackFriend(db *gorm.DB, userId, friendId int64) error {
	var err error
	db = db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	status, err := searchFriendStatus(db, userId, friendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return v1.ErrorErrorFail("无效操作")
		} else {
			return v1.ErrorErrorFail("操作失败")
		}
	}
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_PASS) {
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", userId, friendId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_BLACK, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", friendId, userId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_BLACK, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
	} else {
		return v1.ErrorErrorFail("无效操作")
	}
	return nil
}

//删除好友
func deleteFriend(db *gorm.DB, userId, friendId int64) error {
	var err error
	db = db.Begin()
	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()
	status, err := searchFriendStatus(db, userId, friendId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return v1.ErrorErrorFail("无效操作")
		} else {
			return v1.ErrorErrorFail("操作失败")
		}
	}
	if *status == int32(friend.FriendStatus_FRIEND_STATUS_PASS) || *status == int32(friend.FriendStatus_FRIEND_STATUS_BLACK) {
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", userId, friendId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_DELETE, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
		if err = db.Model(&biz.UserFriend{}).Where("user_id = ? and friend_id = ? and delete_time is null", friendId, userId).
			Updates(map[string]interface{}{"friend_status": friend.FriendStatus_FRIEND_STATUS_DELETE, "update_time": time.Now()}).Error; err != nil {
			return v1.ErrorErrorFail("操作失败")
		}
	} else {
		return v1.ErrorErrorFail("无效操作")
	}
	return nil
}

//查找好友关系
func searchFriendStatus(db *gorm.DB, userId, friendId int64) (*int32, error) {
	var status int32
	if err := db.Model(&biz.UserFriend{}).Select("friend_status").Where("user_id = ? and friend_id = ? and delete_time is null",
		userId, friendId).First(&status).Error; err != nil {
		return nil, err
	}
	return &status, nil
}

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
	"github.com/ilovesusu/suim/app/user/service/internal/pkg"
	"time"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

// CreateUser 创建用户
func (uc *UserRepo) CreateUser(ctx context.Context, user *biz.UserInfo) error {
	user.CreateTime = time.Now()
	var (
		db     = uc.data.db.Model(&biz.UserInfo{})
		number string
		err    error
	)
	//创建用户号码
	for {
		number, err = pkg.RandNumber(uc.log)
		if err != nil {
			return v1.ErrorFAIL("创建失败")
		}
		var con int64
		if err = db.Select("number").Where("number = ? and delete_time is null", number).Count(&con).Error; err != nil {
			uc.log.Log(log.LevelError, "查询用户号码失败", err)
			return v1.ErrorFAIL("创建失败")
		}
		if con == 0 {
			break
		}
	}
	user.Number = number
	//查询手机号是否被使用
	var con int64
	if err = db.Select("phone").Where("delete_time is null").Count(&con).Error; err != nil {
		uc.log.Log(log.LevelError, "查询手机号失败", err)
		return v1.ErrorFAIL("创建用户失败")
	}
	if con > 0 {
		return v1.ErrorPhoneIsUsed("手机号已被使用")
	}
	//创建用户
	if err = db.Create(user).Error; err != nil {
		uc.log.Log(log.LevelError, "创建用户失败", err)
		return v1.ErrorFAIL("创建用户失败")
	}
	return nil
}

// UpdateIdCard 修改身份信息
func (uc *UserRepo) UpdateIdCard(ctx context.Context, req *biz.UpdateIdCardReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"name": req.Name, "id_card": req.IdCard}).
		Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户身份信息失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdatePhone 修改手机号
func (uc *UserRepo) UpdatePhone(ctx context.Context, req *biz.UpdatePhoneReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		con int64
		err error
	)
	if err = db.Select("phone").Where("delete_time is null").Count(&con).Error; err != nil {
		uc.log.Log(log.LevelError, "查询电话号码失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	if con > 0 {
		return v1.ErrorPhoneIsUsed("手机号码已被使用")
	}
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"phone": req.Phone}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户电话号码失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdatePassword 修改密码
func (uc *UserRepo) UpdatePassword(ctx context.Context, req *biz.UpdatePasswordReq) error {
	var (
		db       = uc.data.db.Model(&biz.UserInfo{})
		password string
		err      error
	)
	if err = db.Select("password").Where("id = ?", req.Id).First(&password).Error; err != nil {
		uc.log.Log(log.LevelError, "查询用户密码失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	if password != req.OldPassword {
		return v1.ErrorFAIL("密码不正确")
	}
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"password": req.NewPassword}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户密码失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// ForgetPassword 忘记密码
func (uc *UserRepo) ForgetPassword(ctx context.Context, req *biz.ForgetPasswordReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("phone = ? and delete_time is null", req.Phone).Updates(map[string]interface{}{"password": req.Password}).Error; err != nil {
		uc.log.Log(log.LevelError, "忘记密码重置失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateNickname 修改昵称
func (uc *UserRepo) UpdateNickname(ctx context.Context, req *biz.UpdateNicknameReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"nickname": req.Nickname}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户昵称失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateSex 修改性别
func (uc *UserRepo) UpdateSex(ctx context.Context, req *biz.UpdateSexReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"sex": req.Sex}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户性别失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateAvatarUrl 修改头像
func (uc *UserRepo) UpdateAvatarUrl(ctx context.Context, req *biz.UpdateAvatarUrlReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"avatar_url": req.AvatarUrl}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户头像失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdatePersonalSign 修改个性签名
func (uc *UserRepo) UpdatePersonalSign(ctx context.Context, req *biz.UpdatePersonalSignReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"personal_sign": req.PersonalSign}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户个性签名失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateIntroduce 修改个人介绍
func (uc *UserRepo) UpdateIntroduce(ctx context.Context, req *biz.UpdateIntroduceReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"introduce": req.Introduce}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改个人介绍失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateSnapCall 修改是否允许临时会话
func (uc *UserRepo) UpdateSnapCall(ctx context.Context, req *biz.UpdateSnapCallReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"snap_call": req.SnapCall}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户是否允许临时会话失败", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// UpdateFriendPass 修改添加好友方式
func (uc *UserRepo) UpdateFriendPass(ctx context.Context, req *biz.UpdateFriendPassReq) error {
	var (
		db  = uc.data.db.Model(&biz.UserInfo{})
		err error
	)
	if err = db.Where("id = ?", req.Id).Updates(map[string]interface{}{"friend_pass_type": req.FriendPassType,
		"friend_pass_problem": req.FriendPassProblem, "friend_pass_answer": req.FriendPassAnswer}).Error; err != nil {
		uc.log.Log(log.LevelError, "修改用户添加好友方式", err)
		return v1.ErrorFAIL("修改失败")
	}
	return nil
}

// DeleteUser 删除账户
func (uc *UserRepo) DeleteUser(ctx context.Context, req *biz.DeleteUserReq) error {
	//todo 删除账户并删除该用户其他的信息:群组,好友,频道...
	panic("implement me")
}

// InfoUserBase 用户基本信息
func (uc *UserRepo) InfoUserBase(ctx context.Context, id int64) (*biz.InfoUserBaseRsp, error) {
	var (
		db   = uc.data.db.Model(&biz.UserInfo{})
		data = &biz.InfoUserBaseRsp{}
		err  error
	)
	if err = db.Select("number", "nickname", "sex", "avatar_url", "personal_sign", "introduce").Where("id = ?", id).
		First(data).Error; err != nil {
		uc.log.Log(log.LevelError, "查询用户基本信息失败", err)
		return nil, v1.ErrorFAIL("查询失败")
	}
	return data, nil
}

// InfoAccount 用户身份信息
func (uc *UserRepo) InfoAccount(ctx context.Context, id int64) (*biz.InfoAccountRsp, error) {
	var (
		db   = uc.data.db.Model(&biz.UserInfo{})
		data = &biz.InfoAccountRsp{}
		err  error
	)
	if err = db.Select("phone", "name", "id_card").Where("id = ?", id).First(data).Error; err != nil {
		uc.log.Log(log.LevelError, "查询用户身份信息失败", err)
		return nil, v1.ErrorFAIL("查询失败")
	}
	return data, nil
}

// InfoSnapCall 用户是否允许临时会话
func (uc *UserRepo) InfoSnapCall(ctx context.Context, id int64) (*bool, error) {
	var (
		db   = uc.data.db.Model(&biz.UserInfo{})
		data *bool
		err  error
	)
	if err = db.Select("snap_call").Where("id = ?", id).First(data).Error; err != nil {
		uc.log.Log(log.LevelError, "查询用户是否允许临时会话失败", err)
		return nil, v1.ErrorFAIL("查询失败")
	}
	return data, nil
}

// InfoFriendPass 用户添加好友方式
func (uc *UserRepo) InfoFriendPass(ctx context.Context, id int64) (*biz.InfoFriendPassRsp, error) {
	var (
		db   = uc.data.db.Model(&biz.UserInfo{})
		data = &biz.InfoFriendPassRsp{}
		err  error
	)
	if err = db.Select("friend_pass_type", "friend_pass_problem", "friend_pass_answer").Where("id = ?", id).
		First(data).Error; err != nil {
		uc.log.Log(log.LevelError, "查询用户添加好友方式失败", err)
		return nil, v1.ErrorFAIL("查询失败")
	}
	return data, nil
}

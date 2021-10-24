package biz

import (
	"context"
	v1 "github.com/ilovesusu/suim/api/user/service/v1"
	"strings"
)

// UserInfo 用户信息
type UserInfo struct {
	BaseModel
	Number            string  `gorm:"index;not null;size:50;comment:用户号码"`
	Phone             string  `gorm:"index;not null;size:20;comment:电话号码"`
	Password          string  `gorm:"not null;size:255;comment:密码"`
	Name              *string `gorm:"index;size:50;comment:姓名"`
	IdCard            *string `gorm:"size:20;comment:身份证号"`
	Nickname          string  `gorm:"index;not null;size:100;comment:昵称"`
	Sex               int32   `gorm:"index;not null;size:4;comment:性别(1-保密,2-男,3-女)"`
	AvatarUrl         *string `gorm:"size:255;comment:头像地址链接"`
	PersonalSign      *string `gorm:"size:50;comment:个性签名"`
	Introduce         *string `gorm:"size:255;comment:个人介绍"`
	SnapCall          bool    `gorm:"not null;comment:是否允许临时会话"`
	FriendPassType    int32   `gorm:"not null;comment:添加好友方式(1-直接通过,2-需要验证,3-回答问题通过验证,4-拒绝加好友)"`
	FriendPassProblem *string `gorm:"size:255;comment:问题通过好友请求问题"`
	FriendPassAnswer  *string `gorm:"size:255;comment:问题通过好友答案"`
}

// CreateUser 创建用户
func (u *UserUsecase) CreateUser(ctx context.Context, user *UserInfo) error {
	if user.FriendPassType == int32(v1.FriendPassType_QUESTIONS_AND_ANSWERS) {
		if user.FriendPassProblem == nil || (user.FriendPassProblem != nil && len(strings.TrimSpace(*user.FriendPassProblem)) == 0) {
			return v1.ErrorQuestionsNotNull("通过问题添加好友,问题不能为空")
		}
		if user.FriendPassAnswer == nil || (user.FriendPassAnswer != nil && len(strings.TrimSpace(*user.FriendPassAnswer)) == 0) {
			return v1.ErrorAnswersNotNull("通过问题添加好友,答案不能为空")
		}
	}
	return u.repo.CreateUser(ctx, user)
}

// UpdateAccount 修改身份信息
func (u *UserUsecase) UpdateAccount(ctx context.Context, req *UpdateIdCardReq) error {
	return u.repo.UpdateIdCard(ctx, req)
}

// UpdatePhone 修改电话号码
func (u *UserUsecase) UpdatePhone(ctx context.Context, req *UpdatePhoneReq) error {
	return u.repo.UpdatePhone(ctx, req)
}

// UpdatePassword 修改密码
func (u *UserUsecase) UpdatePassword(ctx context.Context, req *UpdatePasswordReq) error {
	return u.repo.UpdatePassword(ctx, req)
}

// ForgetPassword 忘记密码
func (u *UserUsecase) ForgetPassword(ctx context.Context, req *ForgetPasswordReq) error {
	//todo 通过nats验证发送的验证码
	return u.repo.ForgetPassword(ctx, req)
}

// UpdateNickname 修改昵称
func (u *UserUsecase) UpdateNickname(ctx context.Context, req *UpdateNicknameReq) error {
	return u.repo.UpdateNickname(ctx, req)
}

// UpdateSex 修改性别
func (u *UserUsecase) UpdateSex(ctx context.Context, req *UpdateSexReq) error {
	return u.repo.UpdateSex(ctx, req)
}

// UpdateAvatarUrl 修改头像
func (u *UserUsecase) UpdateAvatarUrl(ctx context.Context, req *UpdateAvatarUrlReq) error {
	return u.repo.UpdateAvatarUrl(ctx, req)
}

// UpdatePersonalSign 修改个性签名
func (u *UserUsecase) UpdatePersonalSign(ctx context.Context, req *UpdatePersonalSignReq) error {
	return u.repo.UpdatePersonalSign(ctx, req)
}

// UpdateIntroduce 修改自我介绍
func (u *UserUsecase) UpdateIntroduce(ctx context.Context, req *UpdateIntroduceReq) error {
	return u.repo.UpdateIntroduce(ctx, req)
}

// UpdateSnapCall 修改是否允许临时会话
func (u *UserUsecase) UpdateSnapCall(ctx context.Context, req *UpdateSnapCallReq) error {
	return u.repo.UpdateSnapCall(ctx, req)
}

// UpdateFriendPass 修改添加好友方式
func (u *UserUsecase) UpdateFriendPass(ctx context.Context, req *UpdateFriendPassReq) error {
	if req.FriendPassType == int32(v1.FriendPassType_QUESTIONS_AND_ANSWERS) {
		if req.FriendPassProblem == nil || (req.FriendPassProblem != nil && len(strings.TrimSpace(*req.FriendPassProblem)) == 0) {
			return v1.ErrorQuestionsNotNull("通过问题添加好友,问题不能为空")
		}
		if req.FriendPassAnswer == nil || (req.FriendPassAnswer != nil && len(strings.TrimSpace(*req.FriendPassAnswer)) == 0) {
			return v1.ErrorAnswersNotNull("通过问题添加好友,答案不能为空")
		}
	}
	return u.repo.UpdateFriendPass(ctx, req)
}

// DeleteUser 删除账户
func (u *UserUsecase) DeleteUser(ctx context.Context, req *DeleteUserReq) error {
	return u.repo.DeleteUser(ctx, req)
}

// InfoUserBase 查看基本信息
func (u *UserUsecase) InfoUserBase(ctx context.Context, id int64) (*InfoUserBaseRsp, error) {
	return u.repo.InfoUserBase(ctx, id)
}

// InfoAccount 查看身份信息
func (u *UserUsecase) InfoAccount(ctx context.Context, id int64) (*InfoAccountRsp, error) {
	return u.repo.InfoAccount(ctx, id)
}

// InfoSnapCall 查看是否允许临时会话
func (u *UserUsecase) InfoSnapCall(ctx context.Context, id int64) (*bool, error) {
	return u.repo.InfoSnapCall(ctx, id)
}

// InfoFriendPass 查看添加好友方式
func (u *UserUsecase) InfoFriendPass(ctx context.Context, id int64) (*InfoFriendPassRsp, error) {
	return u.repo.InfoFriendPass(ctx, id)
}

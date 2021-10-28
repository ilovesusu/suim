package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUsecase, NewFriendUsecase)

// NewUserUsecase 新增用户使用方案
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// NewFriendUsecase 新增好友使用方案
func NewFriendUsecase(repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{repo: repo, log: log.NewHelper(logger)}
}

// UserRepo 用户回购接口
type UserRepo interface {
	// CreateUser 创建用户
	CreateUser(ctx context.Context, user *UserInfo) error
	// UpdateIdCard 修改用户身份信息
	UpdateIdCard(ctx context.Context, req *UpdateIdCardReq) error
	// UpdatePhone 修改电话号码
	UpdatePhone(ctx context.Context, req *UpdatePhoneReq) error
	// UpdatePassword 修改密码
	UpdatePassword(ctx context.Context, req *UpdatePasswordReq) error
	// ForgetPassword 忘记密码
	ForgetPassword(ctx context.Context, req *ForgetPasswordReq) error
	// UpdateNickname 修改昵称
	UpdateNickname(ctx context.Context, req *UpdateNicknameReq) error
	// UpdateSex 修改性别
	UpdateSex(ctx context.Context, req *UpdateSexReq) error
	// UpdateAvatarUrl 修改头像
	UpdateAvatarUrl(ctx context.Context, req *UpdateAvatarUrlReq) error
	// UpdatePersonalSign 修改个性签名
	UpdatePersonalSign(ctx context.Context, req *UpdatePersonalSignReq) error
	// UpdateIntroduce 修改自我介绍
	UpdateIntroduce(ctx context.Context, req *UpdateIntroduceReq) error
	// UpdateSnapCall 修改是否允许临时会话
	UpdateSnapCall(ctx context.Context, req *UpdateSnapCallReq) error
	// UpdateFriendPass 修改用户添加好友方式
	UpdateFriendPass(ctx context.Context, req *UpdateFriendPassReq) error
	// DeleteUser 删除账户
	DeleteUser(ctx context.Context, req *DeleteUserReq) error
	// InfoUserBase 查询用户基本信息
	InfoUserBase(ctx context.Context, id int64) (*InfoUserBaseRsp, error)
	// InfoAccount 查询用户帐号信息
	InfoAccount(ctx context.Context, id int64) (*InfoAccountRsp, error)
	// InfoSnapCall 查询用户是否允许临时会话
	InfoSnapCall(ctx context.Context, id int64) (*bool, error)
	// InfoFriendPass 查询用户添加好友方式
	InfoFriendPass(ctx context.Context, id int64) (*InfoFriendPassRsp, error)
}

// UserUsecase 用户使用方案结构体
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// FriendRepo 好友回购接口
type FriendRepo interface {
	// ListUserFriend 好友列表
	ListUserFriend(ctx context.Context, req *ListFriendReq) ([]ListUserFriendRsp, error)
	// UpdateFriendRemark 修改好友备注
	UpdateFriendRemark(ctx context.Context, req *UpdateFriendRemarkReq) error
	// UpdateFriendStatus 修改好友状态
	UpdateFriendStatus(ctx context.Context, req *UpdateFriendStatusReq) error
}

// FriendUsecase 好友使用方案结构体
type FriendUsecase struct {
	repo FriendRepo
	log  *log.Helper
}

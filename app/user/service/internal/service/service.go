package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
	"github.com/ilovesusu/suim/api/user/service/v1/user"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService, NewFriendService)

// UserService 用户服务
type UserService struct {
	user.UnsafeUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService 创建一个用户服务
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

// FriendService 好友服务
type FriendService struct {
	friend.UnsafeFriendServer

	uc  *biz.FriendUsecase
	log *log.Helper
}

// NewFriendService 创建一个用户服务
func NewFriendService(uc *biz.FriendUsecase, logger log.Logger) *FriendService {
	return &FriendService{uc: uc, log: log.NewHelper(logger)}
}

package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUsecase, NewFriendUsecase)

// NewUserUsecase 新增用户使用方案
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// NewFriendUsecase .
func NewFriendUsecase(repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{repo: repo, log: log.NewHelper(logger)}
}

package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

// UserService 用户服务
type UserService struct {
	v1.UnimplementedUserServer

	uc  *biz.UserUsecase
	log *log.Helper
}

// NewUserService 创建一个用户服务
func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{uc: uc, log: log.NewHelper(logger)}
}

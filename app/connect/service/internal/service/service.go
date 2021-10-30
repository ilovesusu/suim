package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/ilovesusu/suim/api/logic/service/v1"
	"github.com/ilovesusu/suim/app/connect/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopAdmin)

type ShopAdmin struct {
	log *log.Helper
	uc  *biz.UserUseCase
}

func NewShopAdmin(uc *biz.UserUseCase, logger log.Logger) *ShopAdmin {
	return &ShopAdmin{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
	}
}
func (s *ShopAdmin) GetUser(ctx context.Context, in *v1.Input) (*v1.Output, error) {
	return s.uc.GetUser(ctx, in)
}

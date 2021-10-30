package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	logic "github.com/ilovesusu/suim/api/logic/service/v1"
	v1 "github.com/ilovesusu/suim/api/logic/service/v1"
)

type User struct {
	Id int64
}

type UserRepo interface {
}

type UserUseCase struct {
	lc  logic.LogicClient
	log *log.Helper
}

func NewUserUseCase(logger log.Logger, lc logic.LogicClient) *UserUseCase {
	logs := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		lc:  lc,
		log: logs,
	}
}

func (s *UserUseCase) GetUser(ctx context.Context, in *v1.Input) (*v1.Output, error) {
	return s.lc.Call(ctx, in)
}

package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v12 "github.com/ilovesusu/suim/api/logic/service/v1"
	"github.com/ilovesusu/suim/app/logic/service/internal/biz"
)

// LogicServer is a greeter service.
type LogicServer struct {
	v12.UnsafeLogicServer

	uc  *biz.LogicUsecase
	log *log.Helper
}

func (l *LogicServer) Call(ctx context.Context, input *v12.Input) (*v12.Output, error) {
	panic("implement me")
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.LogicUsecase, logger log.Logger) *LogicServer {
	return &LogicServer{uc: uc, log: log.NewHelper(logger)}
}

package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/log"
	v12 "github.com/ilovesusu/suim/api/user/service/v1"
	"github.com/ilovesusu/suim/app/logic/service/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v12.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements user.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v12.HelloRequest) (*v12.HelloReply, error) {
	s.log.WithContext(ctx).Infof("SayHello Received: %v", in.GetName())
	p := &biz.Greeter{}
	err := s.uc.Delete(ctx, p)
	if err != nil {
		return nil, err
	}
	jsonCodec := encoding.GetCodec("json")
	marshal, err := jsonCodec.Marshal(&p)
	if in.GetName() == "error" {
		return nil, v12.ErrorUserNotFound("user not found: %s", in.GetName())
	}
	return &v12.HelloReply{Message: "Hello " + in.GetName() + string(marshal)}, nil
}

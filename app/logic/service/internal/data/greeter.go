package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/app/logic/service/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.Logic {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz.Greeter) error {
	return nil
}

func (r *greeterRepo) DeleteGreeter(ctx context.Context, greeter *biz.Greeter) error {
	panic("implement me")
}

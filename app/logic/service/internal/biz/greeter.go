package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type Logic interface {
	CreateGreeter(context.Context, *Greeter) error
	UpdateGreeter(context.Context, *Greeter) error
	DeleteGreeter(context.Context, *Greeter) error
}

type LogicUsecase struct {
	repo Logic
	log  *log.Helper
}

func NewLogicUsecase(repo Logic, logger log.Logger) *LogicUsecase {
	return &LogicUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LogicUsecase) Create(ctx context.Context, g *Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *LogicUsecase) Update(ctx context.Context, g *Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}

func (uc *LogicUsecase) Delete(ctx context.Context, g *Greeter) error {
	return uc.repo.DeleteGreeter(ctx, g)
}

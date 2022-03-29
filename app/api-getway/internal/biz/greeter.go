package biz

import (
	"context"

	v1 "api-getway/api/getway/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Getway is a Getway model.
type Getway struct {
	Hello string
}

type Goods struct {
	GoodsID   int64
	GoodsName string
	Price     float32
}

// GetwayRepo is a Greater repo.
type GetwayRepo interface {
	Save(context.Context, *Getway) (*Getway, error)
	Update(context.Context, *Getway) (*Getway, error)
	FindByID(context.Context, int64) (*Getway, error)
	ListByHello(context.Context, string) ([]*Getway, error)
	ListAll(context.Context) ([]*Getway, error)

	FindGoodsByID(context.Context, int64) (*Goods, error)
	ListAllGoods(context.Context) ([]*Goods, error)
}

// GetwayUsecase is a Getway usecase.
type GetwayUsecase struct {
	repo GetwayRepo
	log  *log.Helper
}

// NewGetwayUsecase new a Getway usecase.
func NewGetwayUsecase(repo GetwayRepo, logger log.Logger) *GetwayUsecase {
	return &GetwayUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGetway creates a Getway, and returns the new Getway.
func (uc *GetwayUsecase) CreateGetway(ctx context.Context, g *Getway) (*Getway, error) {
	uc.log.WithContext(ctx).Infof("CreateGetway: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

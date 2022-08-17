package biz

import (
	"context"

	v1 "learn-k8s/api/goods/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrGoodsNotFound is goods not found.
	ErrGoodsNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "goods not found")
)

// Goods is a Goods model.
type Goods struct {
	ID    int64
	Price float32
	Name  string
	Sign  string
}

// GoodsRepo is a Goods repo.
type GoodsRepo interface {
	Save(context.Context, *Goods) (*Goods, error)
	Update(context.Context, *Goods) (*Goods, error)
	FindByID(context.Context, int64) (*Goods, error)
	ListByHello(context.Context, string) ([]*Goods, error)
	ListAll(context.Context) ([]*Goods, error)
}

// GreeterUsecase is a Greeter usecase.
type GoodsUsecase struct {
	repo GoodsRepo
	log  *log.Helper
}

// NewGoodsUsecase new a Goods usecase.
func NewGoodsUsecase(repo GoodsRepo, logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GoodsUsecase) GetGoods(ctx context.Context, id int64) (*Goods, error) {
	return uc.repo.FindByID(ctx, id)
}

// CreateGoods creates a Goods, and returns the new Goods.
func (uc *GoodsUsecase) CreateGoods(ctx context.Context, g *Goods) (*Goods, error) {
	// uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GoodsUsecase) ListAll(ctx context.Context) ([]*Goods, error) {
	items, err := uc.repo.ListAll(ctx)
	if err != nil {
		uc.log.WithContext(ctx).Warn()
	}
	return items, err
}

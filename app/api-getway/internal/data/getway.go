package data

import (
	"context"

	"api-getway/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type getwayRepo struct {
	data *Data
	log  *log.Helper
}

// NewGetwayRepo .
func NewGetwayRepo(data *Data, logger log.Logger) biz.GetwayRepo {
	return &getwayRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *getwayRepo) Save(ctx context.Context, g *biz.Getway) (*biz.Getway, error) {
	return g, nil
}

func (r *getwayRepo) Update(ctx context.Context, g *biz.Getway) (*biz.Getway, error) {
	return g, nil
}

func (r *getwayRepo) FindByID(ctx context.Context, id int64) (*biz.Getway, error) {
	return nil, nil
}

func (r *getwayRepo) ListByHello(context.Context, string) ([]*biz.Getway, error) {
	return nil, nil
}

func (r *getwayRepo) ListAll(context.Context) ([]*biz.Getway, error) {
	return nil, nil
}

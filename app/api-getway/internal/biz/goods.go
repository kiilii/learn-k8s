package biz

import "context"

// GetGoods get goods by id
func (uc *GetwayUsecase) GetGoods(ctx context.Context, id int64) (*Goods, error) {
	return uc.repo.FindGoodsByID(ctx, id)
}

func (uc *GetwayUsecase) ListAllGoods(ctx context.Context) ([]*Goods, error) {
	return uc.repo.ListAllGoods(ctx)
}

package data

import (
	"context"
	"sync"

	"learn-k8s/app/goods/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

func faker() *Data {
	var sign = uuid.NewString()
	var faker = map[int64]*biz.Goods{
		1: {ID: 1, Name: "豆浆", Price: 3, Sign: sign},
		2: {ID: 2, Name: "油条", Price: 2, Sign: sign},
		3: {ID: 3, Name: "热干面", Price: 4.5, Sign: sign},
		4: {ID: 4, Name: "杂粮煎饼", Price: 5, Sign: sign},
		5: {ID: 5, Name: "酱饼", Price: 5, Sign: sign},
	}

	return &Data{m: faker, count: int64(len(faker)), mu: new(sync.RWMutex), sign: sign}
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *goodsRepo) Save(ctx context.Context, one *biz.Goods) (*biz.Goods, error) {
	r.data.mu.Lock()
	defer r.data.mu.Unlock()

	r.data.count = r.data.count + 1
	one.ID = int64(r.data.count)
	r.data.m[one.ID] = one

	return one, nil
}

func (r *goodsRepo) Update(ctx context.Context, g *biz.Goods) (*biz.Goods, error) {
	r.data.mu.RLock()
	defer r.data.mu.RUnlock()

	r.data.m[g.ID] = g

	return g, nil
}

func (r *goodsRepo) FindByID(ctx context.Context, id int64) (*biz.Goods, error) {
	r.data.mu.RLock()
	defer r.data.mu.RUnlock()

	if one, has := r.data.m[id]; !has {
		r.log.WithContext(ctx).Errorf("[find]: %v", biz.ErrGoodsNotFound)
		return nil, biz.ErrGoodsNotFound
	} else {
		return one, nil
	}
}

func (r *goodsRepo) ListByHello(context.Context, string) ([]*biz.Goods, error) {
	return nil, nil
}

func (r *goodsRepo) ListAll(context.Context) ([]*biz.Goods, error) {
	r.data.mu.RLock()
	defer r.data.mu.RUnlock()

	var list = make([]*biz.Goods, 0)

	for _, item := range r.data.m {
		list = append(list, &biz.Goods{
			ID:    item.ID,
			Name:  item.Name,
			Price: item.Price,
			Sign:  item.Sign,
		})
	}

	return list, nil
}

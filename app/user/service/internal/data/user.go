package data

import (
	"context"
	"encoding/json"
	"fmt"

	"learn-k8s/app/user/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

const prefix = "user:"
const key = "users"

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx context.Context, one *biz.User) (*biz.User, error) {
	return one, r.data.kv.HSet(ctx, key, one.ID, marshal(one)).Err()
}

func (r *userRepo) Update(ctx context.Context, one *biz.User) (*biz.User, error) {
	return one, r.data.kv.HSet(ctx, key, one.ID, marshal(one)).Err()
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*biz.User, error) {
	value := r.data.kv.HGet(ctx, key, fmt.Sprint(id))
	return unmarshal(value)
}

func (r *userRepo) ListByHello(context.Context, string) ([]*biz.User, error) {
	return nil, nil
}

func (r *userRepo) ListAll(context.Context) ([]*biz.User, error) {
	return nil, nil
}

func unmarshal(cmd *redis.StringCmd) (*biz.User, error) {
	var ret = new(biz.User)

	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	byt, _ := cmd.Bytes()
	err := json.Unmarshal(byt, ret)

	return ret, err
}

func marshal(one *biz.User) []byte {
	ret, _ := json.Marshal(one)
	return ret
}

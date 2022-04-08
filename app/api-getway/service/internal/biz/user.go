package biz

import (
	"context"
	v1 "learn-k8s/api/user/v1"
)

type User struct {
	ID      int64
	Gender  int32
	Name    string
	Desc    string
	Phone   string
	Address string
}

func (uc *GetwayUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	one, err := uc.user.GetUser(ctx, &v1.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return &User{
		ID:      one.GetId(),
		Gender:  one.GetGender(),
		Name:    one.GetName(),
		Desc:    one.GetDesc(),
		Phone:   one.GetPhone(),
		Address: one.GetAddress(),
	}, nil
}

func (uc *GetwayUsecase) GetUsers(ctx context.Context) ([]*User, error) {
	reply, err := uc.user.ListUser(ctx, &v1.ListUserRequest{})
	if err != nil {
		return nil, err
	}

	var items = reply.GetList()
	var list []*User = make([]*User, len(items))
	for i, item := range reply.GetList() {
		list[i] = &User{
			ID:      item.GetId(),
			Gender:  item.GetGender(),
			Name:    item.GetName(),
			Desc:    item.GetDesc(),
			Phone:   item.GetPhone(),
			Address: item.GetAddress(),
		}
	}
	return list, nil
}

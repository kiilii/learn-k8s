package service

import (
	"context"
	v1 "learn-k8s/api/user/v1"
)

func (gs *GetwayService) GetUser(ctx context.Context, in *v1.GetUserRequest) (*v1.GetUserReply, error) {
	one, err := gs.uc.GetUser(ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserReply{
		Id:      one.ID,
		Name:    one.Name,
		Gender:  one.Gender,
		Desc:    one.Desc,
		Phone:   one.Phone,
		Address: one.Address,
	}, err
}

func (gs *GetwayService) GetUsers(ctx context.Context, in *v1.ListUserRequest) (*v1.ListUserReply, error) {
	items, err := gs.uc.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	var list []*v1.Person = make([]*v1.Person, len(items))
	for i, item := range items {
		list[i] = &v1.Person{
			Id:      item.ID,
			Name:    item.Name,
			Gender:  item.Gender,
			Desc:    item.Desc,
			Phone:   item.Phone,
			Address: item.Address,
		}
	}

	return &v1.ListUserReply{List: list}, err
}

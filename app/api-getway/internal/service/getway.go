package service

import (
	"context"

	v1 "api-getway/api/getway/v1"
	"api-getway/internal/biz"
)

// SayHello implements helloworld.GetwayServer.
func (s *GetwayService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.CreateGetway(ctx, &biz.Getway{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}

// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type GoodsHTTPServer interface {
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsReply, error)
	ListGoods(context.Context, *ListGoodsRequest) (*ListGoodsReply, error)
}

func RegisterGoodsHTTPServer(s *http.Server, srv GoodsHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/goods/{id}", _Goods_GetGoods1_HTTP_Handler(srv))
	r.GET("/v1/goods", _Goods_ListGoods1_HTTP_Handler(srv))
}

func _Goods_GetGoods1_HTTP_Handler(srv GoodsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.goods.v1.Goods/GetGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetGoods(ctx, req.(*GetGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetGoodsReply)
		return ctx.Result(200, reply)
	}
}

func _Goods_ListGoods1_HTTP_Handler(srv GoodsHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListGoodsRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.goods.v1.Goods/ListGoods")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListGoods(ctx, req.(*ListGoodsRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListGoodsReply)
		return ctx.Result(200, reply)
	}
}

type GoodsHTTPClient interface {
	GetGoods(ctx context.Context, req *GetGoodsRequest, opts ...http.CallOption) (rsp *GetGoodsReply, err error)
	ListGoods(ctx context.Context, req *ListGoodsRequest, opts ...http.CallOption) (rsp *ListGoodsReply, err error)
}

type GoodsHTTPClientImpl struct {
	cc *http.Client
}

func NewGoodsHTTPClient(client *http.Client) GoodsHTTPClient {
	return &GoodsHTTPClientImpl{client}
}

func (c *GoodsHTTPClientImpl) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...http.CallOption) (*GetGoodsReply, error) {
	var out GetGoodsReply
	pattern := "/v1/goods/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.goods.v1.Goods/GetGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *GoodsHTTPClientImpl) ListGoods(ctx context.Context, in *ListGoodsRequest, opts ...http.CallOption) (*ListGoodsReply, error) {
	var out ListGoodsReply
	pattern := "/v1/goods"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.goods.v1.Goods/ListGoods"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

// Code generated by Kitex v0.4.4. DO NOT EDIT.

package interactservice

import (
	"context"
	interact "github.com/Pinklr/tiktok_demo/kitex_gen/interact"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Favorite(ctx context.Context, Req *interact.FavoriteRequest, callOptions ...callopt.Option) (r *interact.FavoriteResponse, err error)
	FavoriteList(ctx context.Context, Req *interact.FavoriteListRequest, callOptions ...callopt.Option) (r *interact.FavoriteListResponse, err error)
	CommentAction(ctx context.Context, Req *interact.CommentActionRequest, callOptions ...callopt.Option) (r *interact.CommentActionResponse, err error)
	CommentList(ctx context.Context, Req *interact.CommentListRequest, callOptions ...callopt.Option) (r *interact.CommentListResponse, err error)
	CountVideoGetFavorite(ctx context.Context, Req *interact.CountVideoGetFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error)
	CountVideoGetComment(ctx context.Context, Req *interact.CountVideoGetCommentRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error)
	CountUserGetFavorite(ctx context.Context, Req *interact.CountUserGetFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error)
	CountUserFavorite(ctx context.Context, Req *interact.CountUserFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kInteractServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kInteractServiceClient struct {
	*kClient
}

func (p *kInteractServiceClient) Favorite(ctx context.Context, Req *interact.FavoriteRequest, callOptions ...callopt.Option) (r *interact.FavoriteResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Favorite(ctx, Req)
}

func (p *kInteractServiceClient) FavoriteList(ctx context.Context, Req *interact.FavoriteListRequest, callOptions ...callopt.Option) (r *interact.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}

func (p *kInteractServiceClient) CommentAction(ctx context.Context, Req *interact.CommentActionRequest, callOptions ...callopt.Option) (r *interact.CommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kInteractServiceClient) CommentList(ctx context.Context, Req *interact.CommentListRequest, callOptions ...callopt.Option) (r *interact.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}

func (p *kInteractServiceClient) CountVideoGetFavorite(ctx context.Context, Req *interact.CountVideoGetFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountVideoGetFavorite(ctx, Req)
}

func (p *kInteractServiceClient) CountVideoGetComment(ctx context.Context, Req *interact.CountVideoGetCommentRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountVideoGetComment(ctx, Req)
}

func (p *kInteractServiceClient) CountUserGetFavorite(ctx context.Context, Req *interact.CountUserGetFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountUserGetFavorite(ctx, Req)
}

func (p *kInteractServiceClient) CountUserFavorite(ctx context.Context, Req *interact.CountUserFavoriteRequest, callOptions ...callopt.Option) (r *interact.CountResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CountUserFavorite(ctx, Req)
}

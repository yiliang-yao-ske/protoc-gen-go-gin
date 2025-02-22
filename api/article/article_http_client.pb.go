// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.2
// - protoc            v4.24.2
// source: api/article.proto

package article

import (
	context "context"
	ecode "github.com/sunmi-OS/gocore/v2/api/ecode"
	http_request "github.com/sunmi-OS/gocore/v2/utils/http-request"
)

// BlogServiceHTTPClient is the client API for BlogService service.
type BlogServiceHTTPClient interface {
	// 获取文章列表
	// 可以读取不多于999个文章列表
	GetArticles(context.Context, *GetArticlesReq) (*TResponse[GetArticlesReply], error)
	// 新建文章
	CreateArticle(context.Context, *Article) (*TResponse[Article], error)
	// 获取文章详情(TODO get方法还未支持)
	GetOneArticle(context.Context, *GetArticlesReq) (*TResponse[GetArticlesReply], error)
}

type BlogServiceHTTPClientImpl struct {
	hh *http_request.HttpClient
}

func NewBlogServiceHTTPClient(hh *http_request.HttpClient) BlogServiceHTTPClient {
	return &BlogServiceHTTPClientImpl{hh: hh}
}

func (c *BlogServiceHTTPClientImpl) GetArticles(ctx context.Context, req *GetArticlesReq) (*TResponse[GetArticlesReply], error) {
	resp := &TResponse[GetArticlesReply]{}
	_, err := c.hh.Client.R().SetContext(ctx).SetBody(req).SetResult(resp).Post("/v1/articles")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(resp.Code, resp.Msg)
	}
	return resp, err
}

func (c *BlogServiceHTTPClientImpl) CreateArticle(ctx context.Context, req *Article) (*TResponse[Article], error) {
	resp := &TResponse[Article]{}
	_, err := c.hh.Client.R().SetContext(ctx).SetBody(req).SetResult(resp).Post("/v1/author/:author_id/articles")
	if err != nil {
		return nil, err
	}
	if resp.Code != 1 {
		err = ecode.NewV2(resp.Code, resp.Msg)
	}
	return resp, err
}

func (c *BlogServiceHTTPClientImpl) GetOneArticle(ctx context.Context, req *GetArticlesReq) (*TResponse[GetArticlesReply], error) {
	// TODO: GET method not support
	return nil, ecode.NewV2(-1, "GET method not support")
}

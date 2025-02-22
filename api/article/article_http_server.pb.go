// Code generated by protoc-gen-go-gin. DO NOT EDIT.
// versions:
// - protoc-gen-go-gin v0.0.2
// - protoc            v4.24.2
// source: api/article.proto

package article

import (
	gin "github.com/gin-gonic/gin"
	api "github.com/sunmi-OS/gocore/v2/api"
)

// BlogServiceHTTPServer is the server API for BlogService service.
type BlogServiceHTTPServer interface {
	// 获取文章列表
	// 可以读取不多于999个文章列表
	GetArticles(*api.Context, *GetArticlesReq) (*GetArticlesReply, error)
	// 新建文章
	CreateArticle(*api.Context, *Article) (*Article, error)
	// 获取文章详情(TODO get方法还未支持)
	GetOneArticle(*api.Context, *GetArticlesReq) (*GetArticlesReply, error)
}

func RegisterBlogServiceHTTPServer(s *gin.Engine, srv BlogServiceHTTPServer) {
	r := s.Group("/")
	r.POST("/v1/author/articles", _BlogService_GetArticles_HTTP_Handler(srv))              // 获取文章列表
	r.POST("/v1/articles", _BlogService_GetArticles_HTTP_Handler(srv))                     // 获取文章列表
	r.POST("/v1/author/:author_id/articles", _BlogService_CreateArticle_HTTP_Handler(srv)) // 新建文章
	r.GET("/v1/get/article", _BlogService_GetOneArticle_HTTP_Handler(srv))                 // 获取文章详情(TODO get方法还未支持)
}

func _BlogService_GetArticles_HTTP_Handler(srv BlogServiceHTTPServer) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := &GetArticlesReq{}
		ctx := api.NewContext(g)
		err := ctx.ShouldBindJSON(req)
		err = checkValidate(err)
		if err != nil {
			setRetJSON(&ctx, nil, err)
			return
		}
		resp, err := srv.GetArticles(&ctx, req)
		setRetJSON(&ctx, resp, err)
	}
}

func _BlogService_CreateArticle_HTTP_Handler(srv BlogServiceHTTPServer) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := &Article{}
		ctx := api.NewContext(g)
		err := ctx.ShouldBindJSON(req)
		err = checkValidate(err)
		if err != nil {
			setRetJSON(&ctx, nil, err)
			return
		}
		resp, err := srv.CreateArticle(&ctx, req)
		setRetJSON(&ctx, resp, err)
	}
}

func _BlogService_GetOneArticle_HTTP_Handler(srv BlogServiceHTTPServer) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := &GetArticlesReq{}
		ctx := api.NewContext(g)
		err := ctx.ShouldBindJSON(req)
		err = checkValidate(err)
		if err != nil {
			setRetJSON(&ctx, nil, err)
			return
		}
		resp, err := srv.GetOneArticle(&ctx, req)
		setRetJSON(&ctx, resp, err)
	}
}

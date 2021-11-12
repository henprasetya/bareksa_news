package news

import (
	"context"

	"github.com/henprasetya/news/pkg/lib/router"
)

func listNews(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.GetListDataNews(id.Id)
		return response{m}, err
	}
}

func listNewsByTopicId(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.GetListDataNewsByTopic(id.Id)
		return response{m}, err
	}
}

func listNewsByTopicDesc(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		str := req.(request_string)
		m, err := s.GetListDataNewsByTopicDesc(str.str)
		return response{m}, err
	}
}

func listNewsByStatus(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		str := req.(request_string)
		m, err := s.GetListDataNewsByStatus(str.str)
		return response{m}, err
	}
}

func postNews(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {

		data := req.(request)

		m, eror := s.PostDataNews(data.News)
		return response{m}, eror
	}
}

func deleteNews(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.DeleteDataNews(id.Id)
		return response{m}, err
	}
}

func HttpHandler(svc Service, r router.Router) {
	handlerList := router.NewHandler(listNews(svc), restDecodeRequestId, restEncodeResponse, restEncodeError)
	handlerListByTopicId := router.NewHandler(listNewsByTopicId(svc), restDecodeRequestId, restEncodeResponse, restEncodeError)
	handlerListByTopicStr := router.NewHandler(listNewsByTopicDesc(svc), restDecodeRequestString, restEncodeResponse, restEncodeError)
	handlerListByStatus := router.NewHandler(listNewsByStatus(svc), restDecodeRequestString, restEncodeResponse, restEncodeError)
	handlerPost := router.NewHandler(postNews(svc), restDecodeRequestBody, restEncodeResponse, restEncodeError)
	handlerDelete := router.NewHandler(deleteNews(svc), restDecodeRequestId, restEncodeResponse, restEncodeError)
	r.GET("/list-news", handlerList)
	r.GET("/list-news/topic", handlerListByTopicId)
	r.GET("/list-news/topic/desc", handlerListByTopicStr)
	r.GET("/list-news/status", handlerListByStatus)
	r.POST("/post-news", handlerPost)
	r.POST("/delete-news", handlerDelete)
}

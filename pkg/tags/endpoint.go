package tags

import (
	"context"

	"github.com/henprasetya/news/pkg/lib/router"
)

func listTags(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.GetListDataTags(id.Id)
		return response{m}, err
	}
}

func postTags(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {

		data := req.(request)

		m, eror := s.PostDataTags(data.Tags)
		return response{m}, eror
	}
}

func deleteTags(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.DeleteDataTags(id.Id)
		return response{m}, err
	}
}

func HttpHandler(svc Service, r router.Router) {
	handlerListTags := router.NewHandler(listTags(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	handlerPostTags := router.NewHandler(postTags(svc), restDecodeRequestBody, restEncodeResponse, restEncodeError)
	handlerDeleteTags := router.NewHandler(deleteTags(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	r.GET("/list-tags", handlerListTags)
	r.POST("/post-tags", handlerPostTags)
	r.POST("/delete-tags", handlerDeleteTags)
}

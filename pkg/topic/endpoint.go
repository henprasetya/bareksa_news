package topic

import (
	"context"

	"github.com/henprasetya/news/pkg/lib/router"
)

func listTopic(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.GetListDataTopic(id.Id)
		return response{m}, err
	}
}

func postTopic(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {

		data := req.(request)

		m, eror := s.PostDataTopic(data.Topic)
		return response{m}, eror
	}
}

func deleteTopic(s Service) router.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		id := req.(request_id)
		m, err := s.DeleteDataTopic(id.Id)
		return response{m}, err
	}
}

func HttpHandler(svc Service, r router.Router) {
	handlerList := router.NewHandler(listTopic(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	handlerPost := router.NewHandler(postTopic(svc), restDecodeRequestBody, restEncodeResponse, restEncodeError)
	handlerDelete := router.NewHandler(deleteTopic(svc), restDecodeRequest, restEncodeResponse, restEncodeError)
	r.GET("/list-topic", handlerList)
	r.POST("/post-topic", handlerPost)
	r.POST("/delete-topic", handlerDelete)
}

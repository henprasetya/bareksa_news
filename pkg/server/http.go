package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/henprasetya/news/pkg/lib/router"
	"github.com/henprasetya/news/pkg/news"
	service "github.com/henprasetya/news/pkg/services"
	"github.com/henprasetya/news/pkg/tags"
	"github.com/henprasetya/news/pkg/topic"
)

type httpServer struct {
	svc   *service.Service
	port  int
	errCh chan error
}

func NewHttpServer(svc *service.Service, port int) Server {
	return &httpServer{
		svc:  svc,
		port: port,
	}
}

func (h *httpServer) Run() {
	log.Printf("start running http server on port %d\n", h.port)

	route := router.NewDefaultRouter()
	topic.HttpHandler(h.svc.TopicService, route)
	news.HttpHandler(h.svc.NewsService, route)
	tags.HttpHandler(h.svc.TagsService, route)

	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), route.Handler())

	if err != nil {
		h.errCh <- err
	}
}

func (h *httpServer) ListenError() <-chan error {
	return h.errCh
}

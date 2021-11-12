package services

import (
	"github.com/henprasetya/news/pkg/lib/repo"
	"github.com/henprasetya/news/pkg/news"
	"github.com/henprasetya/news/pkg/tags"
	"github.com/henprasetya/news/pkg/topic"
)

type Service struct {
	TopicService topic.Service
	NewsService  news.Service
	TagsService  tags.Service
}

func CreateService(r *repo.Repository) *Service {
	topic := topic.NewService(r.TopicRepo)
	news := news.NewService(r.NewsRepo)
	tag := tags.NewService(r.TagsRepo)
	return &Service{
		TopicService: topic,
		NewsService:  news,
		TagsService:  tag,
	}
}

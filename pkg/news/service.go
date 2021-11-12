package news

import (
	"github.com/henprasetya/news/pkg/model"
	repo "github.com/henprasetya/news/pkg/repo/news"
)

type Service interface {
	GetListDataNews(id int64) (*model.Response, error)
	GetListDataNewsByTopic(id int64) (*model.Response, error)
	GetListDataNewsByTopicDesc(desc string) (*model.Response, error)
	GetListDataNewsByStatus(status string) (*model.Response, error)
	PostDataNews(m model.News) (*model.Response, error)
	DeleteDataNews(id int64) (*model.Response, error)
}

type service struct {
	db repo.NewsOperate
}

func NewService(db repo.NewsOperate) Service {
	var svc Service
	svc = &service{
		db: db,
	}
	return svc
}

func (s *service) GetListDataNews(id int64) (*model.Response, error) {
	return s.db.SelectNewsList(id)
}

func (s *service) GetListDataNewsByTopic(id int64) (*model.Response, error) {
	return s.db.SelectNewsListByTopic(id)
}

func (s *service) GetListDataNewsByTopicDesc(desc string) (*model.Response, error) {
	return s.db.SelectNewsListByTopicDesc(desc)
}

func (s *service) GetListDataNewsByStatus(status string) (*model.Response, error) {
	return s.db.SelectNewsListByStatus(status)
}

func (s *service) PostDataNews(m model.News) (*model.Response, error) {
	return s.db.CreateOrUpdateNews(m)
}

func (s *service) DeleteDataNews(id int64) (*model.Response, error) {
	return s.db.DeleteNews(id)
}

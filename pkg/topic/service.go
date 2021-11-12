package topic

import (
	"github.com/henprasetya/news/pkg/model"
	repo "github.com/henprasetya/news/pkg/repo/topic"
)

type Service interface {
	GetListDataTopic(id int64) (*model.Response, error)
	PostDataTopic(m model.Topic) (*model.Response, error)
	DeleteDataTopic(id int64) (*model.Response, error)
}

type service struct {
	db repo.TopicOperate
}

func NewService(db repo.TopicOperate) Service {
	var svc Service
	svc = &service{
		db: db,
	}
	return svc
}

func (s *service) GetListDataTopic(id int64) (*model.Response, error) {
	return s.db.SelectTopicList(id)
}

func (s *service) PostDataTopic(m model.Topic) (*model.Response, error) {
	return s.db.CreateOrUpdateTopic(m)
}

func (s *service) DeleteDataTopic(id int64) (*model.Response, error) {
	return s.db.DeleteTopic(id)
}

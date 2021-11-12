package tags

import (
	"github.com/henprasetya/news/pkg/model"
	repo "github.com/henprasetya/news/pkg/repo/tags"
)

type Service interface {
	GetListDataTags(id int64) (*model.Response, error)
	PostDataTags(model.Tags) (*model.Response, error)
	DeleteDataTags(id int64) (*model.Response, error)
}

type service struct {
	db repo.TagsOperate
}

func NewService(db repo.TagsOperate) Service {
	var svc Service
	svc = &service{
		db: db,
	}
	return svc
}

func (s *service) GetListDataTags(id int64) (*model.Response, error) {
	return s.db.SelectTagsList(id)
}

func (s *service) PostDataTags(tag model.Tags) (*model.Response, error) {
	return s.db.CreateOrUpdateTags(tag)
}

func (s *service) DeleteDataTags(id int64) (*model.Response, error) {
	return s.db.DeleteTags(id)
}

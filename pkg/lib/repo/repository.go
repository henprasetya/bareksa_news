package repo

import (
	"github.com/henprasetya/news/pkg/lib/mysql"
	"github.com/henprasetya/news/pkg/lib/redis"
	"github.com/henprasetya/news/pkg/repo/news"
	"github.com/henprasetya/news/pkg/repo/tags"
	"github.com/henprasetya/news/pkg/repo/topic"
)

type Repository struct {
	TopicRepo topic.TopicOperate
	NewsRepo  news.NewsOperate
	TagsRepo  tags.TagsOperate
}

func CreateRepository() *Repository {
	var mysql = mysql.NewMySql()
	var redis = redis.NewRedis()
	return &Repository{
		TopicRepo: topic.NewTopicData(mysql, redis.Redis),
		NewsRepo:  news.NewNewsData(mysql),
		TagsRepo:  tags.NewTagsData(mysql),
	}
}

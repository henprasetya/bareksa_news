package news

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/henprasetya/news/pkg/lib/mysql"
	"github.com/henprasetya/news/pkg/model"
)

type NewsOperate interface {
	SelectNewsList(id int64) (*model.Response, error)
	SelectNewsListByStatus(status string) (*model.Response, error)
	SelectNewsListByTopic(id int64) (*model.Response, error)
	SelectNewsListByTopicDesc(desc string) (*model.Response, error)
	CreateOrUpdateNews(m model.News) (*model.Response, error)
	DeleteNews(id int64) (*model.Response, error)
}

type newsDb struct {
	db *sql.DB
}

func NewNewsData(db *mysql.Mysql) NewsOperate {
	return &newsDb{
		db: db.DB,
	}
}

func (t *newsDb) SelectNewsList(id int64) (*model.Response, error) {
	var where string
	if id != 0 {
		where = fmt.Sprintf("where id=%d", id)
	}
	query := "select * from news " + where + " order by id asc"
	rows, err := t.db.Query(query)
	defer rows.Close()
	var result *model.Response
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.News
	for rows.Next() {
		var m model.News
		err = rows.Scan(&m.Id, &m.Description, &m.Status, &m.IdTopic)
		data = append(data, &m)
	}
	result = &model.Response{
		Success: true,
		Message: "success",
		Data:    data,
	}
	return result, nil
}

func (t *newsDb) SelectNewsListByStatus(status string) (*model.Response, error) {
	query := "select * from news where status = ? order by id asc"
	rows, err := t.db.Query(query, status)
	defer rows.Close()
	var result *model.Response
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.News
	for rows.Next() {
		var m model.News
		err = rows.Scan(&m.Id, &m.Description, &m.Status, &m.IdTopic)
		data = append(data, &m)
	}
	result = &model.Response{
		Success: true,
		Message: "success",
		Data:    data,
	}
	return result, nil
}

func (t *newsDb) SelectNewsListByTopic(id int64) (*model.Response, error) {
	query := "select * from news where topic_id = ? order by id asc"
	rows, err := t.db.Query(query, id)
	defer rows.Close()
	var result *model.Response
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.News
	for rows.Next() {
		var m model.News
		err = rows.Scan(&m.Id, &m.Description, &m.Status, &m.IdTopic)
		data = append(data, &m)
	}
	result = &model.Response{
		Success: true,
		Message: "success",
		Data:    data,
	}
	return result, nil
}

func (t *newsDb) SelectNewsListByTopicDesc(desc string) (*model.Response, error) {
	query := "select n.* from news n join topic t on t.id = n.topic_id where t.description like ? order by id asc"
	log.Print(query)
	rows, err := t.db.Query(query, "%"+desc+"%")
	defer rows.Close()
	var result *model.Response
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.News
	for rows.Next() {
		var m model.News
		err = rows.Scan(&m.Id, &m.Description, &m.Status, &m.IdTopic)
		data = append(data, &m)
	}
	result = &model.Response{
		Success: true,
		Message: "success",
		Data:    data,
	}
	return result, nil
}

func (t *newsDb) CreateOrUpdateNews(m model.News) (*model.Response, error) {
	update := false
	if m.Id != 0 {
		update = true
	}

	if update {
		query := "update news set description=?, status=?, topic_id=? where id=?"
		t.db.Exec(query, m.Description, m.Status, m.IdTopic, m.Id)

	} else {
		query := "insert into news(description, status, topic_id)values(?, ?, ?)"
		log.Print("INSERT :", m.IdTopic)
		t.db.Exec(query, m.Description, m.Status, m.IdTopic)
	}
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

func (t *newsDb) DeleteNews(id int64) (*model.Response, error) {
	query := "update news set status = 'deleted' where id=?"
	t.db.Exec(query, id)
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

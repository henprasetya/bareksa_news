package tags

import (
	"database/sql"
	"log"

	"github.com/henprasetya/news/pkg/lib/mysql"
	"github.com/henprasetya/news/pkg/model"
)

type TagsOperate interface {
	SelectTagsList(id int64) (*model.Response, error)
	CreateOrUpdateTags(model.Tags) (*model.Response, error)
	DeleteTags(id int64) (*model.Response, error)
}

type tagsDb struct {
	db *sql.DB
}

func NewTagsData(db *mysql.Mysql) TagsOperate {
	return &tagsDb{
		db: db.DB,
	}
}

func (t *tagsDb) SelectTagsList(id int64) (*model.Response, error) {
	//var where string
	//if id != 0 {
	//	where = fmt.Sprintf("where id=%d", id)
	//}
	query := "select * from tags where id = ?"
	// log.Print(query)
	rows, err := t.db.Query(query, id)
	defer rows.Close()
	var result *model.Response
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.Tags
	for rows.Next() {
		var tag model.Tags
		err = rows.Scan(&tag.Id, &tag.Code, &tag.Description, &tag.IdNews)
		// log.Print("data ", tag.Code)
		data = append(data, &tag)
	}
	result = &model.Response{
		Success: true,
		Message: "success",
		Data:    data,
	}
	return result, nil
}
func (t *tagsDb) CreateOrUpdateTags(tag model.Tags) (*model.Response, error) {
	update := false
	if tag.Id != 0 {
		update = true
	}

	if update {
		query := "update tags set code =?, description=?, news_id=? where id=?"
		t.db.Exec(query, tag.Code, tag.Description, tag.IdNews, tag.Id)

	} else {
		query := "insert into tags(code,description,news_id)values(?,?,?)"
		t.db.Exec(query, tag.Code, tag.Description, tag.IdNews)
	}
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

func (t *tagsDb) DeleteTags(id int64) (*model.Response, error) {
	query := "delete from tags where id=?"
	t.db.Exec(query, id)
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

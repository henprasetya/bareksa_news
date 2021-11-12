package topic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/henprasetya/news/pkg/lib/mysql"
	"github.com/henprasetya/news/pkg/model"
)

type RedisEncode struct {
	Param interface{}
}

type TopicOperate interface {
	SelectTopicList(id int64) (*model.Response, error)
	CreateOrUpdateTopic(model.Topic) (*model.Response, error)
	DeleteTopic(id int64) (*model.Response, error)
}

type topicDb struct {
	db    *sql.DB
	redis *redis.Client
}

func NewTopicData(db *mysql.Mysql, redis *redis.Client) TopicOperate {
	return &topicDb{
		db:    db.DB,
		redis: redis,
	}
}

func (t *topicDb) SelectTopicList(id int64) (*model.Response, error) {

	var dataRedis string
	var err error
	if id != 0 {
		dataRedis, err = t.redis.Get(fmt.Sprintf("data_topic_%d", id)).Result()
		if err == redis.Nil {
			return getResponse(true, "success", t.getDataFromDb(id))
		} else {
			log.Print("DATA REDIS:", dataRedis)
			var data model.Topic
			json.Unmarshal([]byte(dataRedis), &data)
			return getResponse(true, "success", data)
		}
	} else {
		dataRedis, err = t.redis.Get("data_topic_all").Result()
		if err == redis.Nil { // unexpected error
			return getResponse(true, "success", t.getDataFromDb(id))
		} else {

			var data []*model.Topic
			json.Unmarshal([]byte(dataRedis), &data)
			return getResponse(true, "success", data)
		}
	}

}

func (t *topicDb) getDataFromDb(id int64) []*model.Topic {
	var where string
	if id != 0 {
		where = fmt.Sprintf("where id=%d", id)
	}
	query := "select * from topic " + where + " order by id asc"
	log.Print(query)
	rows, err := t.db.Query(query)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var data []*model.Topic
	for rows.Next() {
		var topic model.Topic
		err = rows.Scan(&topic.Id, &topic.Description)
		data = append(data, &topic)
	}
	key := "data_topic_all"
	if id != 0 {
		key = fmt.Sprintf("data_topic_%d", id)
		dec, err := json.Marshal(data[0])
		if err != nil {
			log.Print("Gagal Create Json")
		}
		eror := t.redis.Set(key, dec, 0).Err()
		if eror != nil {
			log.Print(eror)
		}
	} else {
		dec, err := json.Marshal(data)
		if err != nil {
			log.Print("Gagal Create Json")
		}
		eror := t.redis.Set(key, dec, 0).Err()
		if eror != nil {
			log.Print(eror)
		}
	}
	return data
}

func (t *topicDb) CreateOrUpdateTopic(m model.Topic) (*model.Response, error) {
	var update bool
	log.Print("TOPIC : ", m.Description)
	if m.Id != 0 {
		update = true
	}
	if update {
		query := "update topic set description=? where id=?"
		t.db.Exec(query, m.Description, m.Id)
	} else {
		query := "insert into topic (description) values (?)"
		t.db.Exec(query, m.Description)
	}
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

func (t *topicDb) DeleteTopic(id int64) (*model.Response, error) {
	query := "delete from topic where id=?"
	t.db.Exec(query, id)
	return &model.Response{
		Success: true,
		Message: "success",
		Data:    nil,
	}, nil
}

func getResponse(success bool, message string, data interface{}) (*model.Response, error) {
	return &model.Response{
		Success: success,
		Message: message,
		Data:    data,
	}, nil
}

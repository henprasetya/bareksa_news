package model

type News struct {
	Id          int64  `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	IdTopic     int64  `json:"topicId"`
}

package model

type Tags struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	IdNews      int64  `json:"newsId"`
}

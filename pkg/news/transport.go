package news

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/henprasetya/news/pkg/model"
)

type request struct {
	model.News
}

type request_id struct {
	Id int64
}

type request_string struct {
	str string
}

type response struct {
	*model.Response
}

func restDecodeRequestId(_ context.Context, r *http.Request) (interface{}, error) {
	idstr := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		id = 0
	}
	return request_id{id}, nil
}

func restDecodeRequestString(_ context.Context, r *http.Request) (interface{}, error) {
	str := r.URL.Query().Get("value")

	return request_string{str}, nil
}

func restDecodeRequestBody(_ context.Context, r *http.Request) (interface{}, error) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var param model.News
	err := dec.Decode(&param)
	if err != nil {

	}
	return request{param}, nil
}

func restEncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	log.Print(response)
	return json.NewEncoder(w).Encode(response)
}

func restEncodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
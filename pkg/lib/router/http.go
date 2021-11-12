package router

import (
	"context"
	"net/http"
)

type handler struct {
	endpoint  Endpoint
	decodeReq DecodeRequestFunc
	encodeRes EncodeResponseFunc
	encodeErr EncodeErrorFunc
}

func NewHandler(e Endpoint, d DecodeRequestFunc, en EncodeResponseFunc, er EncodeErrorFunc) *handler {
	h := &handler{
		endpoint:  e,
		decodeReq: d,
		encodeRes: en,
		encodeErr: er,
	}
	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	// log.Printf("REQUEST")

	request, err := h.decodeReq(ctx, r)
	if err != nil {
		h.encodeErr(ctx, err, w)
		return
	}

	response, err := h.endpoint(ctx, request)
	if err != nil {
		h.encodeErr(ctx, err, w)
		return
	}

	err = h.encodeRes(ctx, w, response)
	if err != nil {
		h.encodeErr(ctx, err, w)
		return
	}
}

type DecodeRequestFunc func(context.Context, *http.Request) (request interface{}, err error)
type EncodeResponseFunc func(context.Context, http.ResponseWriter, interface{}) error
type EncodeErrorFunc func(context.Context, error, http.ResponseWriter)

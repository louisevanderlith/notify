package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Publish struct {
}

func (req *Publish) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}

func (req *Publish) Create(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}

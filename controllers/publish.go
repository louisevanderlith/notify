package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Publish struct {
}

func (req *Publish) Post(ctx context.Contexer) (int, interface{}) {
	return http.StatusNotImplemented, nil
}

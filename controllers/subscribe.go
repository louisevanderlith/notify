package controllers

import (
	"log"
	"net/http"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/notify/core"
)

type Subscribe struct {
}

func (req *Subscribe) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusMethodNotAllowed, nil
}

func (req *Subscribe) Create(ctx context.Requester) (int, interface{}) {
	subsc := webpush.Subscription{}
	err := ctx.Body(&subsc)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = core.AddSubscriber(subsc)

	if err != nil {
		log.Println(err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, "Subscibed"
}

package controllers

import (
	"log"
	"net/http"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/louisevanderlith/droxolite/xontrols"
	"github.com/louisevanderlith/notify/core"
)

type Subscribe struct {
	xontrols.APICtrl
}

func (req *Subscribe) Post() {
	subsc := webpush.Subscription{}
	err := req.Body(&subsc)

	if err != nil {
		req.Serve(http.StatusBadRequest, err, nil)
		return
	}

	err = core.AddSubscriber(subsc)

	if err != nil {
		log.Println(err)
		req.Serve(http.StatusInternalServerError, err, nil)
		return
	}

	req.Serve(http.StatusOK, nil, "Subscibed")
}

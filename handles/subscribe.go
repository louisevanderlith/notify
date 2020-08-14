package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/louisevanderlith/notify/core"
)

/*
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}*/

func Subscribe(w http.ResponseWriter, r *http.Request) {
	subsc := webpush.Subscription{}
	err := drx.JSONBody(r, &subsc)

	if err != nil {
		log.Println("Bind Error", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = core.AddSubscriber(subsc)

	if err != nil {
		log.Println("Add Subscriber Error", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = mix.Write(w, mix.JSON("Subscribed"))

	if err != nil {
		log.Println("Serve Error", err)
	}
}

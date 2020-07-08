package handles

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Publish(w http.ResponseWriter, r *http.Request) {
	return http.StatusNotImplemented, nil
}

func Subscribe(ctx context.Requester) (int, interface{}) {
	return http.StatusNotImplemented, nil
}

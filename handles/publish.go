package handles

import (
	"net/http"
)

func Publish(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "", http.StatusNotImplemented)
}

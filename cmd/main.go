package main

import (
	"flag"
	"github.com/louisevanderlith/notify/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/notify/core"
)

func main() {
	security := flag.String("security", "http://localhost:8086", "Security Provider's URL")
	srcSecrt := flag.String("scopekey", "secret", "Secret used to validate against scopes")
	flag.Parse()

	core.CreateContext()
	defer core.Shutdown()

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8103",
		Handler:      handles.SetupRoutes(*srcSecrt, *security),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

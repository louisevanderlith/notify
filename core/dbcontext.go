package core

import (
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/serials"
)

type context struct {
	Subscribers husk.Tabler
}

var ctx context

func CreateContext() {

	ctx = context{
		Subscribers: husk.NewTable(Subscriber{}, serials.GobSerial{}),
	}
}

func Shutdown() {
	ctx.Subscribers.Save()
}

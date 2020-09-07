package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Subscribers husk.Table
}

var ctx context

func CreateContext() {

	ctx = context{
		Subscribers: husk.NewTable(Subscriber{}),
	}
}

func Shutdown() {
	ctx.Subscribers.Save()
}
